package consumer

import (
	"log"

	"github.com/streadway/amqp"
)

func ConsumeEvents(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"sensor_events", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()
}
