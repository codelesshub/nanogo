package main

import (
	"github.com/codelesshub/nanogo/config/env"
	"github.com/codelesshub/nanogo/config/rabbitmq"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	body := map[string]interface{}{
		"message": "Hello, world!",
		"foo":     "bar",
	}

	rabbitmq.Publish(exchange().Name, queue().Key, body)

}

func exchange() rabbitmq.Exchange {
	return rabbitmq.Exchange{
		Name:    "teste-exchange-go",
		Durable: true,
		Type:    "direct",
		AutoDel: false,
		NoWait:  false,
	}
}

func queue() rabbitmq.Queue {
	return rabbitmq.Queue{
		Name:       "teste-queue",
		Durable:    true,
		AutoDel:    false,
		Exclusive:  false,
		NoWait:     false,
		Parameters: nil,
	}
}
