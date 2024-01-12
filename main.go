package main // Declare the main package

import (
  "net/http"   // Import the net/http package for HTTP functionality

  "github.com/gin-gonic/gin" // Import the Gin framework for web development
)

func main() {
  // Create a new Gin router
  router := gin.Default()

  // Define a GET route for "/ping"
  router.GET("/ping", func(c *gin.Context) {
    // Send a JSON response with the message "JAISHREERAM" and a 200 OK status code
    c.JSON(http.StatusOK, gin.H{
      "message": "JAISHREERAM",
    })
  })

  // Start the Gin server and listen on the default port (8080)
  router.Run()
}
