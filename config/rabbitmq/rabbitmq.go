package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/codelesshub/nanogo/config/env"
	logger "github.com/codelesshub/nanogo/config/log"

	"github.com/streadway/amqp"
)

type Connection struct {
	*amqp.Connection
	Channel *amqp.Channel
}

var instance *Connection
var once sync.Once

func NewInstanceRabbitmq() *Connection {
	once.Do(func() {
		rabbitmqUser := env.GetEnv("RABBITMQ_USER")
		rabbitmqPassword := env.GetEnv("RABBITMQ_PASSWORD")
		rabbitmqHost := env.GetEnv("RABBITMQ_HOST")
		rabbitmqPort := env.GetEnv("RABBITMQ_PORT")
		rabbitmqVhost := env.GetEnv("RABBITMQ_VHOST")

		url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort, rabbitmqVhost)

		conn, err := amqp.Dial(url)

		if err != nil {
			logger.Fatal("Failed to connect to RabbitMQ: %s", err)
		}

		ch, err := conn.Channel()

		if err != nil {
			logger.Fatal("Failed to open a channel: %s", err)
		}

		instance = &Connection{conn, ch}
	})

	return instance
}
