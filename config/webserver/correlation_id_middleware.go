package webserver

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/codelesshub/nanogo/config/log"
)

func CorrelationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
			r.Header.Set("X-Correlation-ID", correlationID)
		}

		// Criando uma nova instância de logrus com o field definido
		logger := log.LoadLog(correlationID)

		// Adicionando o logger ao contexto da requisição
		ctx := context.WithValue(r.Context(), "logger", logger)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
