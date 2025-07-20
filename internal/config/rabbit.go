package config

import (
	"github.com/streadway/amqp"
	"log"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	_, err = ch.QueueDeclare(
		"sensor_events", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	log.Println("RabbitMQ connected and queue declared")
	return conn, ch
}
