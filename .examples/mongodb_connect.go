package main

import (
	mongodb "github.com/codelesshub/nanogo/config/database"
	"github.com/codelesshub/nanogo/config/env"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Connecta no mongodb
	mongodb.ConnectMongoDB()
}
