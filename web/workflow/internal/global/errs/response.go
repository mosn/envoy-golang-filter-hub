package errs

import (
	"envoy-go-fliter-hub/internal/global/logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type responseBody struct {
	Code   int32  `json:"code"`
	Msg    string `json:"msg"`
	Origin string `json:"origin,omitempty"`
	Data   any    `json:"data,omitempty"`
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
		logs.NameSpace("errs.Fail").Error(fmt.Sprintf("%+v", err))
		e = serverInternal.WithOrigin(err)
	}
	//// If it is a server error, the error details are masked
	////if e.Code/100 == 500 && config.Get().RunMode == config.ModeRelease {
	////	response.Code = serverInternal.Code
	////	response.Msg = serverInternal.Message
	////} else {
	response.Code = e.Code
	response.Msg = e.Message
	response.Origin = e.Origin
	////}

	c.JSON(int(e.Code/100), response)
	c.Abort()
}

func Recovery(c *gin.Context) {
	info := recover()
	if info != nil {
		err, ok := info.(error)

		if ok {
			Fail(c, errors.WithStack(err))
		} else {
			Fail(c, errors.New(fmt.Sprintf("%+v", info)))
		}
	}
	return
}
