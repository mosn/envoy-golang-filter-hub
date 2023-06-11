package controller

type IPluginController interface {
	PluginGet(PluginGetRequest) (*PluginGetResponse, error)
	PluginList(PluginListRequest) (*PluginListResponse, error)
}
