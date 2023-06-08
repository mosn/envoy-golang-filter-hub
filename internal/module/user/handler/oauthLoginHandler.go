package handler

import (
	"envoy-golang-filter-hub/internal/global/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h UserHandler) OauthLogin(c *gin.Context) {
	authURL := oauth.GitHub.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}
