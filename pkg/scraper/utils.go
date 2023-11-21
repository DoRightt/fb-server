package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	internalErr "projects/fb-server/errors"
	"projects/fb-server/pkg/cfg"
	"projects/fb-server/pkg/httplib"
	"projects/fb-server/pkg/model"
	"projects/fb-server/pkg/pgxs"
	fighterRepo "projects/fb-server/repo/fighters"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func ReadFighterData() ([]model.Fighter, error) {
	filePath := "data/fighters.json"

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var fightersData struct {
		Fighters []model.Fighter `json:"Fighters"`
	}

	if err := json.Unmarshal(jsonData, &fightersData); err != nil {
		return nil, err
	}

	return fightersData.Fighters, nil
}

func WriteFighterData(ctx context.Context, l *zap.SugaredLogger, data []model.Fighter) error {
	db, err := pgxs.NewPool(ctx, l, cfg.ViperPostgres())
	if err != nil {
		l.Fatalf("Unable to connect postgresql: %s", err)
	}
	defer db.Pool.Close()

	rep := fighterRepo.New(db)
	counter := 1

	for _, fighter := range data {
		fighterId, err := rep.FindFighter(ctx, fighter)
		if err != nil {
			if err == pgx.ErrNoRows {
				if err := createFighterTx(ctx, rep, fighter); err != nil {
					l.Errorf("Error while fighter transaction: %s", err)
					return err
				}

				fmt.Printf("[Operation №%d] Created: %s\n", counter, fighter.Name)
			} else {
				l.Errorf("Failed to find fighter: %s", err)
				return err
			}
		} else {
			fighterReq := model.FighterReq{
				FighterId: fighterId,
				Fighter:   &fighter,
			}

			if err := updateFighterTx(ctx, rep, fighterReq); err != nil {
				l.Errorf("Error while fighter transaction: %s", err)
				return err
			}

			fmt.Printf("[Operation №%d] Updated: %s\n", counter, fighter.Name)
		}
		counter += 1
	}

	return nil
}

func createFighterTx(ctx context.Context, rep *fighterRepo.FighterRepo, fighter model.Fighter) error {
	tx, err := rep.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		l.Errorf("Unable to begin transaction: %s", err)
	}
	defer tx.Rollback(ctx)

	fighterId, err := rep.CreateNewFighter(ctx, tx, fighter)
	if err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			l.Errorf("Unable to rollback transaction: %s", txErr)
		}
		if err.(*pgconn.PgError).Code == pgerrcode.UniqueViolation {
			intErr := internalErr.New(internalErr.TxNotUnique)
			return httplib.NewApiErrFromInternalErr(intErr)
		} else {
			intErr := internalErr.New(internalErr.TxUnknown)
			l.Errorf("Failed to create fighter during registration transaction: %s", err)
			return httplib.NewApiErrFromInternalErr(intErr, http.StatusInternalServerError)
		}
	}
	stats := model.FighterStatsReq{
		FighterStats: &fighter.Stats,
		FighterId:    fighterId,
	}
	rep.CreateNewFighterStats(ctx, tx, stats)

	if txErr := tx.Commit(ctx); txErr != nil {
		l.Errorf("Unable to commit transaction: %s", txErr)
		return txErr
	}

	return nil
}

func updateFighterTx(ctx context.Context, rep *fighterRepo.FighterRepo, fighter model.FighterReq) error {
	tx, err := rep.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		l.Errorf("Unable to begin transaction: %s", err)
	}
	defer tx.Rollback(ctx)

	updatedId, err := rep.UpdateFighter(ctx, tx, fighter)
	if err != nil {
		if txErr := tx.Rollback(ctx); txErr != nil {
			l.Errorf("Unable to rollback transaction: %s", txErr)
		}
		if err.(*pgconn.PgError).Code == pgerrcode.UniqueViolation {
			intErr := internalErr.New(internalErr.TxNotUnique)
			return httplib.NewApiErrFromInternalErr(intErr)
		} else {
			intErr := internalErr.New(internalErr.TxUnknown)
			l.Errorf("Failed to update fighter during registration transaction: %s", err)
			return httplib.NewApiErrFromInternalErr(intErr, http.StatusInternalServerError)
		}
	}

	stats := model.FighterStatsReq{
		FighterStats: &fighter.Stats,
		FighterId:    updatedId,
	}

	if err := rep.UpdateFighterStats(ctx, tx, stats); err != nil {
		l.Errorf("Error while updating stats: %s", err)
	}

	if txErr := tx.Commit(ctx); txErr != nil {
		l.Errorf("Unable to commit transaction: %s", txErr)
		return txErr
	}

	return nil
}