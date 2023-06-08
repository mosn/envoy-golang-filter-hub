package controller

import (
	"envoy-golang-filter-hub/internal/module/user/handler/dto"
)

type IUserController interface {
	Ping(req dto.PingRequest) (dto.PingReply, error)
}

type UserController struct {
}

func NewUserController() IUserController {
	return UserController{}
}
