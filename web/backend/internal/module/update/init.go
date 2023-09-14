package update

import (
	"envoy-go-fliter-hub/config"
	"github.com/go-git/go-git/v5"
)

type Config struct {
	RemoteRepoUrl string
	LocalRepoPath string
}

type IUpdate interface {
	Update() (*git.Repository, error)
}

var Update IUpdate

type update struct {
	config Config
}

func Init() {
	Update = newUpdate(Config{
		RemoteRepoUrl: config.Config.Repo.OriginRepoUrl,
		LocalRepoPath: config.Config.Repo.LocalRepoPath,
	})
}

func newUpdate(config Config) IUpdate {
	return update{config: config}
}
