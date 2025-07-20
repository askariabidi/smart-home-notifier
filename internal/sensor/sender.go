package sensor

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// SensorEvent represents a message sent by a simulated sensor
type SensorEvent struct {
	Sensor    string `json:"sensor"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

// SendSensorEvent publishes a simulated sensor event to RabbitMQ
func SendSensorEvent(ch *amqp.Channel, sensorType, value string) {
	event := SensorEvent{
		Sensor:    sensorType,
		Value:     value,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling event: %s", err)
		return
	}

	err = ch.Publish(
		"",              // exchange
		"sensor_events", // routing key (queue name)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Failed to publish message: %s", err)
	} else {
		log.Printf("Sensor event published: %s", string(body))
	}
}
