package env

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro carregando arquivo .env")
	}

	log.Info("Carregamento do arquivo .env realizado")
}
