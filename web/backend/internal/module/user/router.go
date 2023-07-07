package user

import (
	"envoy-golang-filter-hub/internal/module/user/handler"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) gin.IRoutes {
	userHandler := handler.NewUserHandler()
	router := r.Group("/user")
	{
		router.GET("/oauth/login", userHandler.OauthLogin)
		router.GET("/oauth/callback", userHandler.OAuthCallback)
	}
	return r
}