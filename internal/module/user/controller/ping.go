package controller

import (
	"envoy-golang-filter-hub/internal/global/log"
	"envoy-golang-filter-hub/internal/module/user/handler/dto"
)

func (c UserController) Ping(req dto.PingRequest) (dto.PingReply, error) {
	log.NameSpace("user.controller.Ping").Info("Ping")
	return dto.PingReply{
		Msg: "Hello World!",
	}, nil
}
