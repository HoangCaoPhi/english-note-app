package main

import (
	"hoangcaophi/english-note-app/src/backend/initialization"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initialization.Run()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
