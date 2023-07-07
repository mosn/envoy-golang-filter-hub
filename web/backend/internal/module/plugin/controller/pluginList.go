package controller

type (
	PluginBasic struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Category    string `json:"category"`
		Owner       string `json:"owner"`
		Tags        []Tag  `json:"tags"`
	}

	Tag struct {
		Name string `json:"name"`
	}
)

type PluginListRequest struct {
	Type     string `json:"type"`
	Category string `json:"category"`
	//Tag      string `json:"tag"` TODO
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type PluginListResponse struct {
	Total   int           `json:"total"`
	Plugins []PluginBasic `json:"plugins"`
}
