package main

import (
	"encoding/json"
	"envoy-go-filter-hub/model"
	"envoy-go-filter-hub/template"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var (
	IndexPath   string
	PluginMap   = make(map[string]template.PluginBasic)
	NewReleases = make([]model.Metadata, 0)
	NewVersions = make([]model.Metadata, 0)
)

func AddVersionToIndex(metadata model.Metadata) {
	fmt.Println("Adding version to index file: ", metadata)

	newPluginBasic := template.PluginBasic{
		PathName:    metadata.PathName,
		Name:        metadata.Name,
		Version:     metadata.Version,
		Category:    metadata.Category,
		Description: metadata.Description,
	}

	PluginMap[metadata.PathName] = newPluginBasic

	pluginDetailPath := filepath.Join(RootPath, "plugins", fmt.Sprintf("%s.json", metadata.PathName))
	pluginDetailFile, err := os.OpenFile(pluginDetailPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	pluginDetail := template.PluginDetail{}
	err = json.NewDecoder(pluginDetailFile).Decode(&pluginDetail)
	if err != nil && err != io.EOF {
		panic(err)
	}
	pluginDetailFile.Close()
	pluginDetail.Version = metadata.Version
	pluginDetail.Name = metadata.Name
	pluginDetail.PathName = metadata.PathName
	pluginDetail.Description = metadata.Description
	pluginDetail.Category = metadata.Category
	pluginDetail.Overview = RenderMarkdown(GetPluginReadme(metadata.PathName))
	pluginDetail.Changelog = RenderMarkdown(GetPluginChangelog(metadata.PathName))
	pluginDetail.Config = RenderMarkdown(GetPluginConfig(metadata.PathName))

	newVersion := template.Version{
		Version:    metadata.Version,
		CreatedAt:  metadata.CreatedAt.Format(time.DateOnly),
		CommitHash: metadata.CommitHash,
		CommitUrl:  fmt.Sprintf("%s/commit/%s", GitHubRepoUrl, metadata.CommitHash),
		Downloads: []template.Download{
			{
				Type: "GitHub",
				Url:  fmt.Sprintf("%s/releases/tag/%s", GitHubRepoUrl, metadata.TagName),
			},
		},
	}

	pluginDetail.Versions = append(pluginDetail.Versions, newVersion)

	pluginDetailBytes, err := json.MarshalIndent(pluginDetail, "", "  ")
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(pluginDetailBytes))

	err = os.WriteFile(pluginDetailPath, pluginDetailBytes, 0644)
	if err != nil {
		panic(err)
	}

}

func SaveIndex() {
	for _, metadata := range NewVersions {
		AddVersionToIndex(metadata)
	}

	pluginList := template.PluginList{}

	pluginList.Plugins = make([]template.PluginBasic, 0, len(PluginMap))
	for _, plugin := range PluginMap {
		pluginList.Plugins = append(pluginList.Plugins, plugin)
	}

	// sort
	sort.Slice(pluginList.Plugins, func(i, j int) bool {
		return pluginList.Plugins[i].Name < pluginList.Plugins[j].Name
	})

	pluginList.TotalCount = len(pluginList.Plugins)

	listBytes, err := json.MarshalIndent(pluginList, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(IndexPath, listBytes, 0644)
}

func GetPluginReadme(pluginName string) string {
	pluginReadmePath := filepath.Join(RootPath, "plugins", pluginName, "readme.md")
	pluginReadmeFile, err := os.Open(pluginReadmePath)
	if err != nil {
		panic(err)
	}
	defer pluginReadmeFile.Close()
	pluginReadmeBytes, err := io.ReadAll(pluginReadmeFile)
	if err != nil {
		panic(err)
	}
	return string(pluginReadmeBytes)
}

func GetPluginChangelog(pluginName string) string {
	pluginChangelogPath := filepath.Join(RootPath, "plugins", pluginName, "changelog.md")
	pluginChangelogFile, err := os.Open(pluginChangelogPath)
	if err != nil {
		panic(err)
	}
	defer pluginChangelogFile.Close()
	pluginChangelogBytes, err := io.ReadAll(pluginChangelogFile)
	if err != nil {
		panic(err)
	}
	return string(pluginChangelogBytes)
}

func GetPluginConfig(pluginName string) string {
	pluginConfigPath := filepath.Join(RootPath, "plugins", pluginName, "config.proto")
	pluginConfigFile, err := os.Open(pluginConfigPath)
	if err != nil {
		panic(err)
	}
	defer pluginConfigFile.Close()
	pluginConfigBytes, err := io.ReadAll(pluginConfigFile)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint("```protobuf\n" + string(pluginConfigBytes) + "\n```\n")
}
