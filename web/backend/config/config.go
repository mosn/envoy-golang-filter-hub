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
}

type Routes struct {
	Frontend   string `envconfig:"FRONTEND"`
	Backend    string `envconfig:"BACKEND"`
	GitHubRepo string `envconfig:"GITHUB_REPO"`
}

type Log struct {
	FilePath string `envconfig:"FILE_PATH"`
}
