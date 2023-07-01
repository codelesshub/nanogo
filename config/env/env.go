package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Erro carregando arquivo .env")
	}

	logrus.Info("Carregamento do arquivo .env realizado")
}

func GetEnv(variable string, default_ ...string) string {
	value := os.Getenv(variable)

	if value == "" {
		if len(default_) > 0 {
			return default_[0]
		}
		logrus.Fatalf("A variavel %s não foi definida no arquivo .env", variable)
	}

	return value
}
