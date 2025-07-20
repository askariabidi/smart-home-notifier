package main

import (
	"github.com/askariabidi/smart-home-notifier/internal/api"
	"github.com/askariabidi/smart-home-notifier/internal/config"
	"github.com/askariabidi/smart-home-notifier/internal/consumer"
	"github.com/askariabidi/smart-home-notifier/internal/storage"
)

func main() {
	// Connect to RabbitMQ
	conn, ch := config.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Initialize SQLite DB and create sensor_events table
	storage.InitDB()

	// Start consuming sensor events in background
	go consumer.ConsumeEvents(ch)

	// Start HTTP server (REST + dashboard + simulator)
	api.StartServer(ch)
}
