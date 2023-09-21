package template

type PluginDetail struct {
	PathName    string    `json:"path_name"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Overview    string    `json:"overview"`
	Config      string    `json:"config"`
	Changelog   string    `json:"changelog"`
	Versions    []Version `json:"versions"`
}

type Version struct {
	Version    string     `json:"version"`
	CreatedAt  string     `json:"created_at"`
	CommitHash string     `json:"commit_hash"`
	CommitUrl  string     `json:"commit_url"`
	Downloads  []Download `json:"downloads"`
}

type Download struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
