package controller

import (
	"envoy-golang-filter-hub/internal/global/log"
)

type PingRequest struct {
	Msg string `json:"msg"`
}

type PingReply struct {
	Msg string `json:"msg"`
}

func (c UserController) Ping(req PingRequest) (PingReply, error) {
	log.NameSpace("user.controller.Ping").Info("Ping")
	return PingReply{
		Msg: "Hello World!",
	}, nil
}
