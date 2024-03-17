package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run("3001"); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
