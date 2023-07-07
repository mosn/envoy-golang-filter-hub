package handler

import (
	"envoy-golang-filter-hub/internal/global/errs"
	"envoy-golang-filter-hub/internal/module/user/controller"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) OAuthCallback(c *gin.Context) {
	var req controller.OAuthCallbackRequest
	var rep *controller.OAuthCallbackResponse

	if err := c.ShouldBind(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithDetails(err.Error()))
		return
	}

	rep, err := controller.NewUserController(c.Request.Context()).OAuthCallback(req)
	if err != nil {
		errs.Fail(c, err)
		return
	}

	errs.Success(c, rep)
}
