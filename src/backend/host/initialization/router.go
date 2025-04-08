package initialization

import (
	"hoangcaophi/english-note-app/src/backend/features/users"
	"hoangcaophi/english-note-app/src/backend/features/wordgroups"
	"hoangcaophi/english-note-app/src/backend/features/words"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	users.InitUserRoute(&router.RouterGroup)
	wordgroups.InitWordGroupRouter(&router.RouterGroup)
	words.InitWordRouter(&router.RouterGroup)

	return router
}
