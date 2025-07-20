package consumer

import (
	"encoding/json"
	"log"

	"github.com/askariabidi/smart-home-notifier/internal/storage"
	"github.com/streadway/amqp"
)

type SensorEvent struct {
	Sensor    string `json:"sensor"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

func ConsumeEvents(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"sensor_events", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received: %s", msg.Body)

			var event SensorEvent
			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Printf("Failed to parse message: %s", err)
				continue
			}

			// Save to DB
			storage.InsertEvent(event.Sensor, event.Value, event.Timestamp)
		}
	}()
}
