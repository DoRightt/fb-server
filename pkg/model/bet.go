package model

// Bet represents users bet properties
type Bet struct {
	BetId     int32 `json:"bet_id"`
	FightId   int32 `json:"fight_id"`
	UserId    int32 `json:"user_id"`
	FighterId int32 `json:"fighter_id"`
}
