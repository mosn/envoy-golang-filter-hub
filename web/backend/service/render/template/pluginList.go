package template

type PluginList struct {
	TotalCount int           `json:"total_count"`
	Plugins    []PluginBasic `json:"plugins"`
}

type PluginBasic struct {
	PathName    string `json:"path_name"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
