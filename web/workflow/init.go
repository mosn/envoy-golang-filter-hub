package main

import (
	"encoding/json"
	"envoy-go-filter-hub/template"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func init() {
	if GitHubToken == "" {
		fmt.Println("Error: Not Found GITHUB_TOKEN")
		os.Exit(1)
	} else if GitHubRepo == "" {
		fmt.Println("Error: Not Found GITHUB_REPOSITORY")
		os.Exit(1)
	} else if GitHubActor == "" {
		fmt.Println("Error: Not Found GITHUB_ACTOR")
		os.Exit(1)
	}

	// 打开仓库
	var err error
	Repo, err = git.PlainOpen(RootPath)
	if err != nil {
		panic(err)
	}

	Head, err = Repo.Head()
	if err != nil {
		panic(err)
	}

	HeadCommit, err = Repo.CommitObject(Head.Hash())
	if err != nil {
		panic(err)
	}

	// 读取 index.json 文件
	pluginListFile, err := os.Open(IndexPath)
	if err != nil {
		panic(err)
	}

	pluginList := template.PluginList{}

	// 解析 json
	err = json.NewDecoder(pluginListFile).Decode(&pluginList)
	if err != nil {
		panic(err)
	}

	for _, plugin := range pluginList.Plugins {
		//fmt.Printf("plugin: %+v\n", plugin)
		PluginMap[plugin.PathName] = plugin
	}

	//fmt.Println(PluginMap)
}
