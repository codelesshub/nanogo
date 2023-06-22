package log

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func GetLoggerFromContext(ctx context.Context) *log.Entry {
	logger := ctx.Value("logger")
	if logger == nil {
		LoadLog()
	}
	return logger.(*log.Entry)
}
