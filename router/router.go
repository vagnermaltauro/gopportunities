package router

import (
	"log"
  "os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
  if err := godotenv.Load(); err != nil {
    log.Printf("Failed to load .env file: %v", err)
  }

  portString := os.Getenv("PORT")

  if portString == "" {
    log.Fatal("PORT is not set. Defaulting to 3001")
  }

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(portString); err != nil {
		log.Printf("Failed to run server: %v", err)
	}

  log.Printf("Server is running on port %v", portString)
}
