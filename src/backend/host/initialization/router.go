package initialization

import (
	"hoangcaophi/english-note-app/src/backend/features/users"
	"hoangcaophi/english-note-app/src/backend/features/wordgroups"
	"hoangcaophi/english-note-app/src/backend/features/words"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	users.InitUserRoute(&router.RouterGroup)
	wordgroups.InitWordGroupRouter(&router.RouterGroup)
	words.InitWordRouter(&router.RouterGroup)

	return router
}
