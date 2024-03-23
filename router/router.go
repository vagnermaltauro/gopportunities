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
		log.Fatal("PORT is not defined in .env file")
	}

	r := gin.Default()

  initializeRoutes(r)

	if err := r.Run(":" + portString); err != nil {
		log.Printf("Failed to run server: %v", err)
	}

	log.Printf("Server is running on port %v", portString)
}
