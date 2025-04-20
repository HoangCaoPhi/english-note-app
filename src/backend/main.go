package main

import (
	"hoangcaophi/english-note-app/src/backend/host/initialization"
	"hoangcaophi/english-note-app/src/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initialization.Run()

	router.GET("/ping", func(c *gin.Context) {
		response.Success(c, gin.H{"message": "pong"})
	})
	router.Run()
}
