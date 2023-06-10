package errs

import (
	"envoy-golang-filter-hub/config"
	"envoy-golang-filter-hub/internal/global/logx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseBody struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data ...any) {
	response := responseBody{
		Code: success.Code,
		Msg:  success.Message,
		Data: nil,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	c.JSON(http.StatusOK, response)
}

func Fail(c *gin.Context, err error) {
	var response responseBody

	e, ok := err.(*Error)
	if !ok {
		logs.NameSpace("errs.Fail").Error(err.Error())
		e = serverInternal.WithDetails(err.Error())
	}

	// If it is a server error, the error details are masked
	if e.Code/100 == 500 && config.Get().RunMode == config.ModeRelease {
		response.Code = serverInternal.Code
		response.Msg = serverInternal.Message
	} else {
		response.Code = e.Code
		response.Msg = e.Message
	}

	c.JSON(int(e.Code/100), response)
}
