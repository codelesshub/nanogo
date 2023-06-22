package main

import (
	"github.com/codelesshub/nanogo/config/env"
	redis "github.com/codelesshub/nanogo/config/redis"
	"github.com/codelesshub/nanogo/config/webserver"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	redis.StartRedisQueue()

	// redis.Enqueue("my_job", nil)
	// redis.RedisQueueConsumer(tasks.NewCurlTask(), "my_job")

	// Inicializa o webserver
	webserver.StartWebServer()
}
