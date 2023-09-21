package module

import (
	"envoy-go-fliter-hub/internal/module/consumer"
	"envoy-go-fliter-hub/internal/module/parse"
	"envoy-go-fliter-hub/internal/module/render"
	"envoy-go-fliter-hub/internal/module/server"
	"envoy-go-fliter-hub/internal/module/update"
)

func Init() {
	update.Init()
	parse.Init()
	render.Init()
	server.Init()
	consumer.Init()
}
