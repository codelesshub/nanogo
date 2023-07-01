package rabbitmq

import (
	"encoding/json"
	"log"
)

type Consumer interface {
	Consume(body map[string]interface{}, headers map[string]interface{})
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
			bodyMap := make(map[string]interface{})

			err := json.Unmarshal(d.Body, &bodyMap)
			if err != nil {
				log.Fatalf("Error decoding JSON: %s", err)
				continue
			}

			headerBytes, err := json.Marshal(d.Headers)
			if err != nil {
				log.Fatalf("Error encoding headers to JSON: %s", err)
				continue
			}

			headersMap := make(map[string]interface{})
			err = json.Unmarshal(headerBytes, &headersMap)
			if err != nil {
				log.Fatalf("Error decoding headers JSON: %s", err)
				continue
			}

			consumer.Consume(bodyMap, headersMap)
		}
	}()
}
