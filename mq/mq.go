package mq

import (
	"github.com/streadway/amqp"
)

type RabbitMqHub struct {
	URL       string
	QueueName string

	conn    *amqp.Connection
	channel *amqp.Channel
}
type Options func(*RabbitMqHub)

// NewRabbitMqHub
func NewRabbitMqHub(cfg RabbitMqHub, opts ...Options) *RabbitMqHub {
	hub := &cfg
	for _, opt := range opts {
		opt(hub)
	}

	return hub
}

// Connect
func (hub *RabbitMqHub) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(hub.URL)
	hub.conn = conn
	return conn, err
}

// close
func (hub *RabbitMqHub) Close() error {
	return hub.conn.Close()
}

// channel
func (hub *RabbitMqHub) Channel() (*amqp.Channel, error) {
	channel, err := hub.conn.Channel()
	hub.channel = channel
	return channel, err
}

// ConnectQueue
func (hub *RabbitMqHub) ConnectQueue(q string) (amqp.Queue, error) {
	queue, err := hub.channel.QueueDeclare(q, true, false, false, false, nil)
	return queue, err
}

// channel
func (hub *RabbitMqHub) CloseChannel() error {
	return hub.channel.Close()
}
