package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func StartServer() {
	// Define HTTP route handlers
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/api/data", DataHandler)

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from:", r.RemoteAddr)
	fmt.Fprintf(w, "Hello World")
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate fetching data from a database or external service
	data := struct {
		Message string `json:"message"`
		Time    string `json:"time"`
	}{
		Message: "Data retrieved successfully",
		Time:    time.Now().Format(time.RFC3339),
	}

	// Serialize data to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
