package controller

import (
	"envoy-golang-filter-hub/internal/global/errs"
	"envoy-golang-filter-hub/internal/global/jwt"
	"envoy-golang-filter-hub/internal/global/oauth"
	"github.com/google/go-github/v53/github"
)

type OAuthCallbackRequest struct {
	Code string `form:"code" binding:"required"`
}

type OAuthCallbackResponse struct {
	//AccessToken string `json:"access_token"`
	//Items any `json:"items"`
	Token string `json:"token"`
}

func (c UserController) OAuthCallback(req OAuthCallbackRequest) (*OAuthCallbackResponse, error) {
	// 使用授权码获取访问令牌
	githubToken, err := oauth.GitHub.Exchange(c.ctx, req.Code)
	if err != nil {
		return nil, errs.Code2GitHubTokenFailed.Warp(err)
	}

	client := github.NewClient(oauth.GitHub.Client(c.ctx, githubToken))

	// 查询个人信息
	userInfo, _, err := client.Users.Get(c.ctx, "")
	if err != nil {
		return nil, errs.GetGitHubUserInfoFailed.Warp(err)
	}

	// 生成JWT
	JwtToken, err := jwt.CreateToken(jwt.Payload{
		GitHubUsername: *userInfo.Login,
		GitHubID:       *userInfo.ID,
		AvatarURL:      *userInfo.AvatarURL,
	})
	if err != nil {
		return nil, errs.GenerateTokenFailed.Warp(err)
	}

	return &OAuthCallbackResponse{
		Token: JwtToken,
	}, nil
}
