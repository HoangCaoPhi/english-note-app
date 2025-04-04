package users

import "github.com/gin-gonic/gin"

func InitUserRoute(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("users")

	userRouter.Use()

	userRepositoryReadImpl := NewUserRepositoryReadImpl()
	userRepositoryWriteImpl := NewUserRepositoryWriteImpl()

	InitUserRepositoryRead(userRepositoryReadImpl)
	InitUserRepositoryWrite(userRepositoryWriteImpl)

	refreshTokenRepositoryReadImpl := NewRefreshTokenRepositoryReadImpl()
	refreshTokenRepositoryWriteImpl := NewRefreshTokenRepositoryWriteImpl()

	InitRefreshTokenRepositoryRead(refreshTokenRepositoryReadImpl)
	InitRefreshTokenRepositoryWrite(refreshTokenRepositoryWriteImpl)

	userImpl := NewUserServiceImpl(
		NewUserRepositoryRead(),
		NewUserRepositoryWrite(),
		NewRefreshTokenRepositoryRead(),
		NewRefreshTokenRepositoryWrite())

	InitUserService(userImpl)

	userController := NewUserController(userService)

	userRouter.POST("/register", userController.Register)
	userRouter.POST("/login", userController.Login)
}
