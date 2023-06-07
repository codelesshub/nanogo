package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LoadLog(correlationID string) *log.Entry {
	if correlationID == "" {
		correlationID = "unknown"
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	if getEnvironment() == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	logger := log.WithFields(log.Fields{
		"app":              getApplicationName(),
		"env":              getEnvironment(),
		"version":          getVersion(),
		"x-correlation-id": correlationID,
	})

	return logger
}

func getApplicationName() string {
	application_name := os.Getenv("SERVER_PORT")

	if application_name == "" {
		log.Fatal("O nome da aplicação não foi definida no arquivo .env")
	}

	return application_name
}

func getEnvironment() string {
	env := os.Getenv("ENV")

	if env == "" {
		log.Fatal("O ambiente não foi definida no arquivo .env")
	}

	return env
}

func getVersion() string {
	version := os.Getenv("VERSION")

	if version == "" {
		log.Fatal("A versão não foi definida no arquivo .env")
	}

	return version
}
