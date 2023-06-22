package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger         *logrus.Entry
	logInitialized bool
)

func LoadLog(correlationID ...string) *logrus.Entry {
	var cid string
	if len(correlationID) > 0 {
		cid = correlationID[0]
	} else {
		cid = ""
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)

	if getEnvironment() == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	logger = logrus.WithFields(logrus.Fields{
		"app":              getApplicationName(),
		"env":              getEnvironment(),
		"version":          getVersion(),
		"x-correlation-id": cid,
	})

	logInitialized = true

	return logger
}

func getApplicationName() string {
	application_name := os.Getenv("SERVER_PORT")

	if application_name == "" {
		logrus.Fatal("O nome da aplicação não foi definida no arquivo .env")
	}

	return application_name
}

func getEnvironment() string {
	env := os.Getenv("ENV")

	if env == "" {
		logrus.Fatal("O ambiente não foi definida no arquivo .env")
	}

	return env
}

func getVersion() string {
	version := os.Getenv("VERSION")

	if version == "" {
		logrus.Fatal("A versão não foi definida no arquivo .env")
	}

	return version
}

func Fatal(args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Fatal(args)
}

func Debug(args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Debug(args)
}

func Info(args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Info(args)
}
