package main

import (
	"encoding/json"
	"envoy-go-filter-hub/template"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
)

func init() {

	RootPath = filepath.Join("../../")
	RootPath, _ = filepath.Abs(RootPath)

	IndexPath = filepath.Join(RootPath, "index.json")

	fmt.Println("RootPath: ", RootPath)

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

	// 打开仓库 - Open repository
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

	//RunCommand("cd " + RootPath)
	RunCommand("git checkout cache")

	// 读取 index.json 文件 - Read index.json
	pluginListFile, err := os.Open(IndexPath)
	if err != nil {
		panic(err)
	}

	pluginList := template.PluginList{}

	// 解析 json - Parse json
	err = json.NewDecoder(pluginListFile).Decode(&pluginList)
	if err != nil {
		panic(err)
	}

	for _, plugin := range pluginList.Plugins {
		//fmt.Printf("plugin: %+v\n", plugin)
		PluginMap[plugin.PathName] = plugin
	}

	RunCommand("git checkout main")

	//fmt.Println(PluginMap)
}
