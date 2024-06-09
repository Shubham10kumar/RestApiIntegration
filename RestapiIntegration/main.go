package main

import (
	"RestapiIntegration/constants"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// Start the HTTP server in a separate goroutine
	go StartServer()

	// Open the browser to access the server's URL
	openBrowser("http://localhost:8080")

	// Keep the main goroutine running
	select {}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}

func StartServer() {
	// Define HTTP route handlers
	http.HandleFunc(constants.HomeRoute, HelloHandler)
	http.HandleFunc(constants.DataRoute, DataHandler)

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
