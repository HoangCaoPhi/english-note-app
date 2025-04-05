package initialization

import (
	"hoangcaophi/english-note-app/src/backend/infrastructure/mongodb"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	loadConfiguration()

	mongodb.InitMongoDb()

	router := InitRouter()
	return router
}
