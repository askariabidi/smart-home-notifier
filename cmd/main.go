package main

import (
	"github.com/askariabidi/smart-home-notifier/internal/api"
	"github.com/askariabidi/smart-home-notifier/internal/config"
	"github.com/askariabidi/smart-home-notifier/internal/consumer"
	"github.com/askariabidi/smart-home-notifier/internal/sensor"
	"github.com/askariabidi/smart-home-notifier/internal/storage"
)

func main() {
	// Connect to RabbitMQ
	conn, ch := config.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Initialize SQLite database
	storage.InitDB()

	// Simulate two sensor events
	sensor.SendSensorEvent(ch, "motion", "detected")
	sensor.SendSensorEvent(ch, "temperature", "27.5°C")

	// Start consumer in background
	go consumer.ConsumeEvents(ch)

	// Start the web server (REST + HTML)
	api.StartServer()
}
