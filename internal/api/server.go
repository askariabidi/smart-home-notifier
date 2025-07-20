package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/askariabidi/smart-home-notifier/internal/storage"
)

type SensorEvent struct {
	ID        int    `json:"id"`
	Sensor    string `json:"sensor"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

func StartServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/logs", logsHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

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
