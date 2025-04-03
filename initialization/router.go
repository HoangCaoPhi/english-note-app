package initialization

import (
	"hoangcaophi/english-note-app/src/backend/features/users"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	users.InitUserRoute(&router.RouterGroup)

	return router
}
