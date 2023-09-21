package model

import "time"

type Metadata struct {
	TagName     string    `yaml:"-"`
	CommitHash  string    `yaml:"-"`
	CreatedAt   time.Time `yaml:"-"`
	Name        string    `yaml:"Name"`
	PathName    string    `yaml:"-"`
	Description string    `yaml:"Description"`
	Category    string    `yaml:"Category"`
	Version     string    `yaml:"Version"`
}
