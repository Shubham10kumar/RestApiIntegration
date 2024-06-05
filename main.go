package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Initialize a Gin router
    r := gin.Default()

    // Define route handlers
    r.GET("/", serveHome)
    r.GET("/about", serveAbout)
    r.GET("/contact", serveContact)

    // Start the server on port 8000
    r.Run("localhost:8000")
}

// Handler for the home route
func serveHome(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to our simple REST API",
    })
}

// Handler for the about route
func serveAbout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "About Us: We provide REST API examples",
    })
}

// Handler for the contact route
func serveContact(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Contact Us at contact@example.com",
    })
}

