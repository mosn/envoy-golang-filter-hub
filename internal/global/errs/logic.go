package errs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code    int32             `json:"code"`
	Message string            `json:"msg"`
	LogInfo map[string]string `json:"-"`
}

func Success(c *gin.Context, data ...any) {
	reply := struct {
		Code int32  `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data,omitempty"`
	}{
		Code: 200,
		Msg:  "Success",
		Data: nil,
	}
	if len(data) > 0 {
		reply.Data = data[0]
	}
	c.JSON(http.StatusOK, reply)
}
