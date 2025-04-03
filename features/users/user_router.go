package users

import "github.com/gin-gonic/gin"

func InitUserRoute(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("users")

	userRouter.Use()

	userRepositoryReadImpl := NewUserRepositoryReadImpl()
	userRepositoryWriteImpl := NewUserRepositoryWriteImpl()

	InitUserRepositoryRead(userRepositoryReadImpl)
	InitUserRepositoryWrite(userRepositoryWriteImpl)

	userImpl := NewUserServiceImpl(NewUserRepositoryRead(), NewUserRepositoryWrite())
	InitUserService(userImpl)

	userController := NewUserController(userService)

	userRouter.POST("/register", userController.Register)
}
