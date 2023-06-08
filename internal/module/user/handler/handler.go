package handler

import (
	"envoy-golang-filter-hub/internal/module/user/controller"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Ping(c *gin.Context)
	OauthLogin(c *gin.Context)
	OauthCallback(c *gin.Context)
}

type UserHandler struct {
	UserController controller.IUserController
}

func NewUserHandler() IUserHandler {
	return UserHandler{
		UserController: controller.NewUserController(),
	}
}
