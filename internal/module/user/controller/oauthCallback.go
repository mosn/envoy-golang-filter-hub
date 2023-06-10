package controller

import (
	"envoy-golang-filter-hub/internal/global/errs"
	"envoy-golang-filter-hub/internal/global/oauth"
)

type OAuthCallbackRequest struct {
	Code string `form:"code" binding:"required"`
}

type OAuthCallbackResponse struct {
	AccessToken string `json:"access_token"`
}

func (c UserController) OAuthCallback(req OAuthCallbackRequest) (*OAuthCallbackResponse, error) {
	// 使用授权码获取访问令牌
	token, err := oauth.GitHub.Exchange(c.ctx, req.Code)
	if err != nil {
		return nil, errs.Code2TokenFailed.WithDetails(err.Error())
	}

	// 使用访问令牌进行后续操作，如获取用户信息等
	// 这里只是简单地输出访问令牌信息
	return &OAuthCallbackResponse{
		AccessToken: token.AccessToken,
	}, nil
}
