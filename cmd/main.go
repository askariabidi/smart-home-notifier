package main

import (
	"github.com/askariabidi/smart-home-notifier/internal/api"
	"github.com/askariabidi/smart-home-notifier/internal/config"
	"github.com/askariabidi/smart-home-notifier/internal/consumer"
	"github.com/askariabidi/smart-home-notifier/internal/sensor"
)

func main() {
	conn, ch := config.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Simulate two events
	sensor.SendSensorEvent(ch, "motion", "detected")
	sensor.SendSensorEvent(ch, "temperature", "27.5°C")

	go consumer.ConsumeEvents(ch)

	api.StartServer()
}
