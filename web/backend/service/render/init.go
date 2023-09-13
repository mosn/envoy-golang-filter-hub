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

func (r render) writeToFile(details []template.PluginDetail, list template.PluginList) error {
	//TODO implement me
	panic("implement me")
}
