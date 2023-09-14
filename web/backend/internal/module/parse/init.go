package parse

import (
	"github.com/go-git/go-git/v5"
)

type IParse interface {
	Parse(repo *git.Repository) ([]Metadata, error)
}

var Parse IParse

type parse struct{}

func Init(config Config) error {
	_ = config
	Parse = parse{}
	return nil
}
