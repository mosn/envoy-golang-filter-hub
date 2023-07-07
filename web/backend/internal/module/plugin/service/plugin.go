package service

import (
	"envoy-golang-filter-hub/internal/module/plugin/store/model"
)

type Plugin interface {
	Create(plugin *model.Plugin) error
	Delete(pluginID uint) error
	Update(pluginID uint, plugin *model.Plugin) error
	Get(pluginID uint) (*model.Plugin, error)
	List() ([]*model.Plugin, error)

	AddTag(pluginID uint, tagID uint) error
	DeleteTag(pluginID uint, tagID uint) error

	AddVersion(pluginID uint, version model.Version) error
	DeleteVersion(pluginID uint, versionID uint) error
}
