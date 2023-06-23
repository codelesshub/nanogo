package main

import (
	"github.com/codelesshub/nanogo/config"
	"github.com/codelesshub/nanogo/config/env"
	redis "github.com/codelesshub/nanogo/config/redis"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	//Inicializa a conexão com redis
	redis.StartRedisQueue()

	//Inicializa o NewCurlTask na fila my_job
	redis.RedisQueueConsumer(tasks.NewCurlTask(), "my_job")

	//Aguarda o sinal de stop para parar a aplicação
	config.WaitSignalStop()
}
