package controller

import (
	"github.com/codelesshub/nanogo/config/log"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context())
	logger.Info("Healthcheck request received")

	w.WriteHeader(http.StatusOK)
}
