package parse

type Metadata struct {
	TagName    string `yaml:"-"`
	CommitHash string `yaml:"-"`
	Name       string `yaml:"Name"`
	PathName   string `yaml:"-"`
	Category   string `yaml:"Category"`
	Version    string `yaml:"Version"`
}
