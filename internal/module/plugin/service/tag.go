package service

import "envoy-golang-filter-hub/internal/module/plugin/store/model"

type Tag interface {
	Create(tag *model.Tag) error
	Delete(tagID uint) error
	Update(tagID uint, tag *model.Tag) error
}
