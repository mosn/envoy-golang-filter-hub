package model

import (
	"gorm.io/gorm"
	"time"
)

type Plugin struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex;comment:插件名称"`
	Description string    `gorm:"type:varchar(100);not null;comment:插件描述"`
	Type        string    `gorm:"type:varchar(100);not null;comment:插件类型"`
	Category    string    `gorm:"type:varchar(100);not null;comment:插件分类"`
	Owner       string    `gorm:"type:varchar(100);not null;comment:插件作者，GitHub 仓库所有者"`
	Repository  string    `gorm:"type:varchar(100);not null;comment:插件仓库"`
	README      string    `gorm:"type:longtext;not null;comment:插件 README"`
	CreatedAt   time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"type:datetime;not null;comment:更新时间"`

	// has many
	Versions []Version `gorm:"foreignKey:PluginID"`

	// many to many
	Tags []Tag `gorm:"many2many:plugin_tags;"`
}

type Version struct {
	gorm.Model
	PluginID uint   `gorm:"type:int(11);not null;comment:插件ID"`
	Name     string `gorm:"type:varchar(100);not null;comment:插件版本"`
	URL      string `gorm:"type:varchar(100);not null;comment:插件下载地址"`
}

type Tag struct {
	gorm.Model

	// many to many
	Plugins []Plugin `gorm:"many2many:plugin_tags;"`

	Name string `gorm:"type:varchar(100);not null;uniqueIndex;comment:标签名称"`
}
