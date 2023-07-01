package rabbitmq

type ExchangeType string

const (
	Direct  ExchangeType = "direct"
	Fanout  ExchangeType = "fanout"
	Topic   ExchangeType = "topic"
	Headers ExchangeType = "headers"
)

type Exchange struct {
	Name     string
	Durable  bool
	Type     ExchangeType
	AutoDel  bool
	Internal bool
	NoWait   bool
}

func DeclareExchange(exchange Exchange) error {
	connection := NewInstanceRabbitmq()

	return connection.Channel.ExchangeDeclare(
		exchange.Name,
		string(exchange.Type),
		exchange.Durable,
		exchange.AutoDel,
		exchange.Internal,
		exchange.NoWait,
		nil,
	)
}
