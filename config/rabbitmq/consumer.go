package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume(headers amqp.Table, body []byte)
}

func Consume(exchange Exchange, queue Queue, consumer Consumer) {
	connection := NewInstanceRabbitmq()

	DeclareExchange(exchange)
	DeclareQueue(queue)

	BindQueue(exchange, queue)

	msgs, err := connection.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			consumer.Consume(d.Headers, d.Body)
		}
	}()
}
