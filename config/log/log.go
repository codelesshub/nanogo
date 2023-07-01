package log

import (
	"os"

	"github.com/codelesshub/nanogo/config/env"
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

	if env.GetEnv("ENV", "dev") == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	logger = logrus.WithFields(logrus.Fields{
		"app":              env.GetEnv("APP_NAME"),
		"env":              env.GetEnv("ENV"),
		"version":          env.GetEnv("VERSION"),
		"x-correlation-id": cid,
	})

	logInitialized = true

	return logger
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

func Debugf(format string, args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Infof(format, args)
}

func Fatalf(format string, args ...interface{}) {
	if !logInitialized {
		LoadLog()
	}

	logger.Fatalf(format, args)
}
