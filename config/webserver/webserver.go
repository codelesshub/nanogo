package webserver

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func StartWebServer() {
	port := getPortWebServer()

	fmt.Printf("Servidor iniciado em localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, WebServerRouter()))
}

func getPortWebServer() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatal("A porta do servidor não foi definida no arquivo .env")
	}

	return port
}
