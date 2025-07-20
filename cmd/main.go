package main

import (
	"log"

	"github.com/askariabidi/smart-home-notifier/internal/config"
	"github.com/askariabidi/smart-home-notifier/internal/consumer"
	"github.com/askariabidi/smart-home-notifier/internal/api"
)

func main() {
	// Connect to RabbitMQ
	conn, ch := config.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Start event consumer
	go consumer.ConsumeEvents(ch)

	// Start HTTP server
	log.Println("Starting server at :8080")
	api.StartServer()
}
