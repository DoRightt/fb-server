package services

import (
	"context"
	"fmt"
	"net/http"
	"projects/fb-server/pkg/httplib"
	"strings"

	"github.com/spf13/viper"
)

var allowedHeaders = []string{
	"Accept",
	"Content-Type",
	"Content-Length",
	"Authorization",
	"X-Requested-With",
	// "X-HTTP-Method-Override",
	"Cookie",
}

var allowedMethods = []string{
	http.MethodOptions,
	http.MethodGet,
	http.MethodPost,
	http.MethodPatch,
	http.MethodPut,
	http.MethodDelete,
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.Router.ServeHTTP(w, r.WithContext(ctx))
}

func (h *ApiHandler) RunHTTPServer(ctx context.Context) error {
	httplib.SetCookieName(viper.GetString("auth.cookie_name"))

	h.Router.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

	srvAddr := viper.GetString("http.addr")
	if len(srvAddr) < 1 || strings.Index(srvAddr, ":") < 0 {
		return fmt.Errorf("'%s' service address not specified", h.ServiceName)
	}

	h.Logger.Infof("Start listen '%s' http: %s", h.ServiceName, srvAddr)
	fmt.Printf("Server is listening at: %s\n", srvAddr)
	return http.ListenAndServe(srvAddr, h)
}
