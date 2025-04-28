package broker

import "user/internal/client/broker/rabbitmq"

type ClientMsgBroker interface {
	Connect() *rabbitmq.RabbitMQ
	Close() error
}
