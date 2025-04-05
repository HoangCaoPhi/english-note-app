package wordgroups

import (
	"hoangcaophi/english-note-app/src/backend/host/middlewares"

	"github.com/gin-gonic/gin"
)

func InitWordGroupRouter(router *gin.RouterGroup) {
	authorized := router.Group("/word-groups").Use(middlewares.Authentication())

	wordGroupRepositoryReadImpl := NewWordGroupRepositoryReadImpl()
	wordGroupRepositoryWriteImpl := NewWordGroupRepositoryWriteImpl()

	InitWordGroupRepositoryRead(wordGroupRepositoryReadImpl)
	InitWordGroupRepositoryWrite(wordGroupRepositoryWriteImpl)

	wordGroupServiceImpl := NewWordGroupServiceImpl(NewReadRepositoryRead(), NewReadRepositoryWrite())
	InitWordGroupService(wordGroupServiceImpl)

	wordGroupController := NewWordGroupController(NewWordGroupService())

	authorized.GET("", wordGroupController.GetWordGroupsByUserId)
	authorized.POST("", wordGroupController.CreateWordGroup)
}
