package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogMiddleware(ctx *gin.Context) {
	// TODO: log each request and response gracefully
	fmt.Println(ctx.Request.Method, ctx.Request.URL, ctx.Request.Proto)

	ctx.Next()

	fmt.Println(ctx.Writer.Status(), ctx.Writer.Size())
}
