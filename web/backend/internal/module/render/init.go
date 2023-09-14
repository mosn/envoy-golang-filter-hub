package render

import (
	"envoy-go-fliter-hub/internal/module/parse"
	template2 "envoy-go-fliter-hub/internal/module/render/template"
)

type Config struct {
	OutPutDir string
}

type IRender interface {
	Render([]parse.Metadata) error
	renderIntoStruct([]parse.Metadata) ([]template2.PluginDetail, template2.PluginList, error)
	writeToFile([]template2.PluginDetail, template2.PluginList) error
}

var Render IRender

type render struct {
	config Config
}

//
//func Init(config Config) error {
//	Render, err := newRender(config)
//	return err
//}

func newRender(config Config) IRender {
	return render{config: config}
}
