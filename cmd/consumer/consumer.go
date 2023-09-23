package consumer

import (
	"fmt"
	"log"
	"os"

	"github.com/pcpratheesh/rabbitmq-go-eda-tutorial/models"
	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

var goroutineMap = make(map[string]consumer)

func CloseAllConsumers() {
	for _, consumer := range goroutineMap {
		consumer.Stop()
	}
}

func GetConsumerList() map[string]consumer {
	return goroutineMap
}

// Create a map to store goroutine handlers and their corresponding done channels
type consumer struct {
	name string
	done chan bool
}

// New consumer
func NewConsumer(name string) *consumer {
	c := consumer{
		name: name,
		done: make(chan bool, 1),
	}
	goroutineMap[name] = c
	return &c
}

// Run consumer
func (c *consumer) Run(consumeChannel <-chan amqp.Delivery, responseChannel chan models.WebsocketDataPayload) {

	log.Printf("Consumer %v ready, PID: %d", c.name, os.Getpid())

	for {
		select {
		case <-c.done:
			log.Println("Consumer stopped.")
			return
		case d, ok := <-consumeChannel:
			if !ok {
				log.Println("Consumer channel closed. Exiting.")
				return
			}

			log.Printf("[%v] Received a message: %s", c.name, d.Body)

			responseChannel <- models.WebsocketDataPayload{
				Name: c.name,
				Type: "consumed-data",
				Data: string(d.Body),
			}
			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message: %s", err)
			} else {
				log.Printf("[%v] Acknowledged message", c.name)
			}
		}
	}
}

// Stop the consumer
func (c *consumer) Stop() {
	fmt.Printf("Closing consumer [%v] \n", c.name)
	c.done <- true
	delete(goroutineMap, c.name)
}

// LaunchConsumer
func LaunchConsumer(name string, socketDataPaylod chan models.WebsocketDataPayload) {
	conn, err := amqp.Dial("amqp://user:user@localhost:5672/my_vhost")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	autoAck, exclusive, noLocal, noWait := false, false, false, false

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		autoAck,
		exclusive,
		noLocal,
		noWait,
		nil,
	)
	handleError(err, "Could not register consumer")

	fmt.Println("connection success")

	consumer := NewConsumer(name)
	consumer.Run(messageChannel, socketDataPaylod)
}
