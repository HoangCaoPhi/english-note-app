package words

import (
	"hoangcaophi/english-note-app/src/backend/host/middlewares"

	"github.com/gin-gonic/gin"
)

func InitWordRouter(router *gin.RouterGroup) {
	authorized := router.Group("/words").Use(middlewares.Authentication())

	wordRepositoryReadImpl := NewWordRepositoryReadImpl()
	wordRepositoryWriteImpl := NewWordRepositoryWriteImpl()

	InitWordRepositoryRead(wordRepositoryReadImpl)
	InitWordRepositoryWrite(wordRepositoryWriteImpl)

	wordServiceImpl := NewWordServiceImpl(NewWordRepositoryRead(), NewWordRepositoryWrite())
	InitWordService(wordServiceImpl)

	wordController := NewWordController(NewWordService())

	authorized.POST("", wordController.CreateWord)
	authorized.GET("", wordController.GetWordsByGroupID)
}
