package initialization

import (
	"hoangcaophi/english-note-app/src/backend/features/users"
	"hoangcaophi/english-note-app/src/backend/features/wordgroups"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	users.InitUserRoute(&router.RouterGroup)
	wordgroups.InitWordGroupRouter(&router.RouterGroup)

	return router
}
