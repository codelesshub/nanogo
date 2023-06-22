package main

import (
	"github.com/codelesshub/nanogo/config/env"
	redis "github.com/codelesshub/nanogo/config/redis"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	redis.StartRedisQueue()

	//Inicializa o NewCurlTask na fila my_job
	redis.RedisQueueConsumer(tasks.NewCurlTask(), "my_job")

}
