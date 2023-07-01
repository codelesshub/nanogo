package rabbitmq

func BindQueue(exchange Exchange, queue Queue) error {
	connection := NewInstanceRabbitmq()

	return connection.Channel.QueueBind(
		queue.Name,
		queue.Key,
		exchange.Name,
		false,
		nil,
	)
}
