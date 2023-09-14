package render

import (
	"encoding/json"
	"envoy-go-fliter-hub/service/render/template"
	"envoy-go-fliter-hub/tools"
	"fmt"
	"path"
)

const (
	PluginListFileName = "index.json"
	PluginDetailDir    = "plugins"
	PluginDetailSuffix = ".json"
)

func (r render) writeToFile(details []template.PluginDetail, list template.PluginList) error {

	listBytes, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println("json.Marshal(list) error", err)
		return err
	}
	err = tools.Write(listBytes, path.Join(r.config.OutPutDir, PluginListFileName))
	if err != nil {
		return err
	}

	// 创建插件详情目录
	err = tools.CreateDir(path.Join(r.config.OutPutDir, PluginDetailDir))
	if err != nil {
		return err
	}

	for _, detail := range details {
		detailBytes, err := json.MarshalIndent(detail, "", "    ")
		if err != nil {
			fmt.Println("json.Marshal(detail) error", err)
			return err
		}
		err = tools.Write(detailBytes, path.Join(r.config.OutPutDir, PluginDetailDir, detail.PathName+PluginDetailSuffix))
		if err != nil {
			panic(err)
			return err
		}
	}

	return nil
}
