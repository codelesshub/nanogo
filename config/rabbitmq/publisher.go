package rabbitmq

import (
	"encoding/json"

	logger "github.com/codelesshub/nanogo/config/log"
	"github.com/streadway/amqp"
)

func Publish(exchangeName string, routingKey string, body map[string]interface{}) {
	connection := NewInstanceRabbitmq()

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		logger.Fatalf("Houve uma falha ao converter a struct para json: %s", err)
	}

	errPublish := connection.Channel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		},
	)

	if errPublish != nil {
		logger.Fatalf("Failed to publish a message: %s", errPublish)
	}
}
