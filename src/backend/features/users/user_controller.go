package users

import (
	"context"
	"hoangcaophi/english-note-app/src/backend/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.BindJSON(&registerRequest); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user := CreateUser(registerRequest.Username,
		registerRequest.Email)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := user.HashPassword(registerRequest.Password); err != nil {
		response.InternalServerError(c, "Could not hash password")
		return
	}

	_, err := u.userService.Register(ctx, user)
	if err != nil {
		response.InternalServerError(c, "Could not create user")
		return
	}

	response.Created(c, gin.H{"message": "User registered successfully"})
}

func (u *UserController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	access_token, refresh_token, err := u.userService.Login(c, loginRequest.Username, loginRequest.Password)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	})
}
