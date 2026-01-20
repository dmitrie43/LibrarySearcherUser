package rabbitmq

import (
	"log"

	"github.com/dmitrie43/LibrarySearcherUser/internal/config"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Listening() {
	rabbitConnection, rabbitChan := newConnect()

	defer rabbitChan.Close()
	defer rabbitConnection.Close()

	channel := make(chan bool)

	go listen(channel, rabbitChan, "CreateUser")
	go listen(channel, rabbitChan, "UpdateUser")

	<-channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func newConnect() (*amqp091.Connection, *amqp091.Channel) {
	cnf := config.MustLoad()

	conn, err := amqp.Dial("amqp://" + cnf.RabbitUser + ":" + cnf.RabbitPassword + "@" + cnf.RabbitHost + ":" + cnf.RabbitPort + "/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}

func listen(channel chan bool, ch *amqp091.Channel, queue string) {
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}

	channel <- true
}
