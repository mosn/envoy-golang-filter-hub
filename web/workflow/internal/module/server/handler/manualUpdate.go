package handler

import (
	"envoy-go-fliter-hub/internal/global/errs"
	"envoy-go-fliter-hub/internal/global/mq"
	"github.com/gin-gonic/gin"
)

func (h handler) ManualUpdate(c *gin.Context) {
	mq.MQ.SendUpdateSignal()
	errs.Success(c)
}
