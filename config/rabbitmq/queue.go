package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Queue struct {
	Name       string
	Key        string
	Durable    bool
	AutoDel    bool
	Exclusive  bool
	NoWait     bool
	Parameters amqp.Table
}

func DeclareQueue(queue Queue) (amqp.Queue, error) {
	connection := NewInstanceRabbitmq()

	return connection.Channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDel,
		queue.Exclusive,
		queue.NoWait,
		queue.Parameters,
	)
}
