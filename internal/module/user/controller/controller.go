package controller

import (
	"context"
)

type IUserController interface {
	OAuthCallback(OAuthCallbackRequest) (*OAuthCallbackResponse, error)
}

type UserController struct {
	ctx context.Context
}

func NewUserController(ctx context.Context) IUserController {
	return UserController{
		ctx: ctx,
	}
}
