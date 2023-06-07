package webserver

import (
	"github.com/codelesshub/nanogo/controller"

	"github.com/gorilla/mux"
)

func WebServerRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(CorrelationIDMiddleware)

	router.HandleFunc("/healthcheck", controller.HealthcheckHandler)

	return router
}
