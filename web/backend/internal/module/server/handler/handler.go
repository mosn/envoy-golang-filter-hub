package handler

import "github.com/gin-gonic/gin"

type IHandler interface {
	PluginList(c *gin.Context)
	PluginGet(c *gin.Context)
	ManualUpdate(c *gin.Context)
}

var Handler = handler{}

type handler struct{}
