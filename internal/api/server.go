package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/askariabidi/smart-home-notifier/internal/sensor"
	"github.com/askariabidi/smart-home-notifier/internal/storage"
	"github.com/streadway/amqp"
)

type SensorEvent struct {
	ID        int    `json:"id"`
	Sensor    string `json:"sensor"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

var ch *amqp.Channel // Global RabbitMQ channel

func StartServer(rabbitChannel *amqp.Channel) {
	ch = rabbitChannel // Assign the RabbitMQ channel

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/logs", logsHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/simulate", simulateHandler)

	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Smart Home Event Notification System API")
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Query("SELECT id, sensor, value, timestamp FROM sensor_events ORDER BY id DESC")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []SensorEvent
	for rows.Next() {
		var e SensorEvent
		if err := rows.Scan(&e.ID, &e.Sensor, &e.Value, &e.Timestamp); err != nil {
			continue
		}
		events = append(events, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT sensor, value, MAX(timestamp) as timestamp
		FROM sensor_events
		GROUP BY sensor
	`
	rows, err := storage.DB.Query(query)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Status struct {
		Sensor    string `json:"sensor"`
		Value     string `json:"value"`
		Timestamp string `json:"timestamp"`
	}

	var status []Status
	for rows.Next() {
		var s Status
		if err := rows.Scan(&s.Sensor, &s.Value, &s.Timestamp); err != nil {
			continue
		}
		status = append(status, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT sensor, value, MAX(timestamp) as timestamp
		FROM sensor_events
		GROUP BY sensor
	`
	rows, err := storage.DB.Query(query)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Status struct {
		Sensor    string
		Value     string
		Timestamp string
	}

	var data []Status
	for rows.Next() {
		var s Status
		if err := rows.Scan(&s.Sensor, &s.Value, &s.Timestamp); err != nil {
			continue
		}
		data = append(data, s)
	}

	tmplPath := filepath.Join("internal", "templates", "dashboard.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sensorType := r.FormValue("sensor")
		value := "detected"
		if sensorType == "temperature" {
			value = "27.5°C"
		}
		sensor.SendSensorEvent(ch, sensorType, value)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// GET request - render simulate page
	tmplPath := filepath.Join("internal", "templates", "simulate.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
