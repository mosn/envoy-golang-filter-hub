package parse

import (
	"envoy-go-fliter-hub/internal/model"
	"github.com/go-git/go-git/v5"
)

type IParse interface {
	Parse(repo *git.Repository) ([]model.Metadata, error)
}

var Parse IParse

type parse struct{}

func Init() {
	Parse = parse{}
}
