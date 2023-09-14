package render

import (
	"envoy-go-fliter-hub/service/parse"
	"envoy-go-fliter-hub/service/render/template"
)

type IRender interface {
	Render([]parse.Metadata) error
	renderIntoStruct([]parse.Metadata) ([]template.PluginDetail, template.PluginList, error)
	writeToFile([]template.PluginDetail, template.PluginList) error
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
