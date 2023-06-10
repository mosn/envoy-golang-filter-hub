package handler

import (
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	OauthLogin(c *gin.Context)
	OAuthCallback(c *gin.Context)
}

type UserHandler struct {
}

func NewUserHandler() IUserHandler {
	return UserHandler{}
}
