package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"toddlebank2/api"
)

func main() {
	// Start the HTTP server in a separate goroutine
	api.StartServer()

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
