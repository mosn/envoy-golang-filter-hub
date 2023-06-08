package handler

import (
	"envoy-golang-filter-hub/internal/global/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h UserHandler) OauthCallback(c *gin.Context) {

	code := c.Query("code")

	// 使用授权码获取访问令牌
	token, err := oauth.GitHub.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 使用访问令牌进行后续操作，如获取用户信息等
	// 这里只是简单地输出访问令牌信息
	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken, "token_type": token.TokenType})
}
