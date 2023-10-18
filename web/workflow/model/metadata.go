package model

import "time"

type Metadata struct {
	TagName     string    `yaml:"-"`
	CommitHash  string    `yaml:"-"`
	CreatedAt   time.Time `yaml:"-"`
	Name        string    `yaml:"name"`
	PathName    string    `yaml:"-"`
	Description string    `yaml:"description"`
	Category    string    `yaml:"category"`
	Version     string    `yaml:"version"`
}
