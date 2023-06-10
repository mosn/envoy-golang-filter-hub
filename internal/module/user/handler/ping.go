package handler

import (
	"envoy-golang-filter-hub/internal/global/errs"
	"envoy-golang-filter-hub/internal/module/user/controller"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) Ping(c *gin.Context) {
	var req controller.PingRequest
	var rep controller.PingReply

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rep, err := h.UserController.Ping(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	errs.Success(c, rep)
}
