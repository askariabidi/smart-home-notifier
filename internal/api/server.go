package api

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Smart Home Event Notification System")
	})

	http.ListenAndServe(":8080", nil)
}
