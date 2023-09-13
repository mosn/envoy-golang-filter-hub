package render

import (
	"envoy-go-fliter-hub/service/parse"
	"envoy-go-fliter-hub/service/render/template"
	"github.com/Masterminds/semver/v3"
	"time"
)

func (r render) renderIntoStruct(metadata []parse.Metadata) ([]template.PluginDetail, template.PluginList, error) {
	// 创建一个 map 来存储聚合的 PluginDetail，key 是 PathName
	pluginDetailMap := make(map[string]*template.PluginDetail)

	// 遍历所有元数据并填充 PluginDetail 和 Version
	for _, meta := range metadata {
		detail, ok := pluginDetailMap[meta.PathName]
		if !ok {
			detail = &template.PluginDetail{
				PathName:    meta.PathName,
				Name:        meta.Name,
				Version:     meta.Version,
				Category:    meta.Category,
				Description: meta.Description,
			}
			pluginDetailMap[meta.PathName] = detail
		}

		//fmt.Println("test")
		//fmt.Println(detail.Version, meta.Version)
		if semver.MustParse(detail.Version).LessThan(semver.MustParse(meta.Version)) {
			//fmt.Println("changed")
			pluginDetailMap[meta.PathName].Version = meta.Version
			pluginDetailMap[meta.PathName].Name = meta.Name
			pluginDetailMap[meta.PathName].Category = meta.Category
			pluginDetailMap[meta.PathName].Description = meta.Description
		}

		version := template.Version{
			Version:    meta.Version,
			CreatedAt:  meta.CreatedAt.Format(time.DateOnly),
			CommitHash: meta.CommitHash,
		}

		detail.Versions = append(detail.Versions, version)
	}

	// 将 map 转换为切片
	var pluginDetails []template.PluginDetail
	for _, detail := range pluginDetailMap {
		pluginDetails = append(pluginDetails, *detail)
	}

	// 创建 PluginList 结构体
	pluginList := template.PluginList{
		TotalCount: len(pluginDetails),
		Plugins: func(details []template.PluginDetail) []template.PluginBasic {
			var plugins []template.PluginBasic
			for _, detail := range details {
				plugin := template.PluginBasic{
					PathName:    detail.PathName,
					Name:        detail.Name,
					Version:     detail.Version,
					Category:    detail.Category,
					Description: detail.Description,
				}
				plugins = append(plugins, plugin)
			}
			return plugins
		}(pluginDetails),
	}

	return pluginDetails, pluginList, nil
}
