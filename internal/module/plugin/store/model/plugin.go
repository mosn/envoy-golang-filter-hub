package model

import "gorm.io/gorm"

type Plugin struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;uniqueIndex;comment:插件名称"`
	Desc string `gorm:"type:varchar(100);not null;comment:插件描述"`
	Type string `gorm:"type:varchar(100);not null;comment:插件类型"`
	Repo string `gorm:"type:varchar(100);not null;comment:插件仓库"`
}

type Version struct {
	gorm.Model
	PluginID uint   `gorm:"type:int(11);not null;comment:插件ID"`
	Version  string `gorm:"type:varchar(100);not null;comment:插件版本"`
	URL      string `gorm:"type:varchar(100);not null;comment:插件下载地址"`
}
