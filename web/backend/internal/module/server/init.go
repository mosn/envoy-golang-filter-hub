package server

import (
	"envoy-go-fliter-hub/internal/module/server/handler"
	"github.com/gin-gonic/gin"
)

func Init() {

}

func InitRouter(r *gin.RouterGroup) gin.IRouter {
	router := r.Group("/")
	{
		router.GET("/pluginList", handler.Handler.PluginList)
		router.GET("/plugin/:id", handler.Handler.PluginGet)
		router.GET("/manual_update", handler.Handler.ManualUpdate)
	}
	return r
}
