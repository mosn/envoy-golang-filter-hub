package controller

type (
	PluginDetail struct {
		Repository string `json:"repository"`
		//README     string    `json:"readme"`
		CreateAt  int64     `json:"create_at"`
		UpdatedAt int64     `json:"updated_at"`
		Versions  []Version `json:"versions"`
	}

	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
)

type PluginGetRequest struct {
	Owner string `url:"owner" binding:"required"`
	Name  string `url:"name" binding:"required"`
}

type PluginGetResponse struct {
	PluginBasic
	PluginDetail
}
