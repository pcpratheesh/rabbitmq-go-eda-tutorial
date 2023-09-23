package producer

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func RunProducer() {
	conn, err := amqp.Dial("amqp://user:user@localhost:5672/my_vhost")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	limit := 10
	var wg sync.WaitGroup
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		number := i
		go func() {
			defer wg.Done()
			err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(fmt.Sprintf("Data %v", number)),
			})

			if err != nil {
				log.Fatalf("Error publishing message: %s", err)
			}
		}()
	}

	wg.Wait()
}
