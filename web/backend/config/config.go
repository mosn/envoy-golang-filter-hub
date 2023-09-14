package config

const (
	DebugMode   = "Debug"
	ReleaseMode = "Release"
)

type config struct {
	Name   string `envconfig:"NAME"`
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Prefix string `envconfig:"PREFIX"`
	Mode   string `envconfig:"Mode"` // 运行模式：Debug or Release

	Log    Log    `envconfig:"LOG" prefix:"LOG_"`
	Routes Routes `envconfig:"ROUTES" prefix:"ROUTES_"`
	Repo   Repo   `envconfig:"REPO" prefix:"REPO_"`
}

type Routes struct {
	Frontend string `envconfig:"FRONTEND"`
	Backend  string `envconfig:"BACKEND"`
}

type Log struct {
	FilePath string `envconfig:"FILE_PATH"`
}

type Repo struct {
	LocalRepoPath   string `envconfig:"LOCAL_REPO_PATH"`
	OriginRepoUrl   string `envconfig:"ORIGIN_REPO_URL"`
	IndexOutPutPath string `envconfig:"INDEX_OUTPUT_PATH"`
}
