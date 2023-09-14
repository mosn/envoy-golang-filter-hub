package render

import (
	"envoy-go-fliter-hub/config"
	"envoy-go-fliter-hub/internal/model"
	"envoy-go-fliter-hub/internal/template"
)

type Config struct {
	OutPutDir string
}

type IRender interface {
	Render([]model.Metadata) error
	renderIntoStruct([]model.Metadata) ([]template.PluginDetail, template.PluginList, error)
	writeToFile([]template.PluginDetail, template.PluginList) error
}

var Render IRender

type render struct {
	config Config
}

func Init() {
	Render = newRender(Config{
		OutPutDir: config.Config.Repo.LocalRepoPath,
	})
}

func newRender(config Config) IRender {
	return render{config: config}
}
