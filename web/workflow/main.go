package main

import (
	"envoy-go-filter-hub/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	RootPath string
)

func main() {
	// 遍历 plugins 目录下的所有插件，读取各自的 metadata.yaml， 检查是否符合规范，得到 []Metadata
	pluginsDir := filepath.Join(RootPath, "plugins")
	pluginDirs, _ := os.ReadDir(pluginsDir)
	var ans []model.Metadata
	for _, f := range pluginDirs {
		if f.IsDir() {
			pluginPath := filepath.Join(pluginsDir, f.Name())

			// 读取 metadata.yaml 文件
			metadataFile, err := os.Open(filepath.Join(pluginPath, "metadata.yaml"))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// 解析 yaml
			var metadata model.Metadata
			err = yaml.NewDecoder(metadataFile).Decode(&metadata)
			metadata.PathName = f.Name()
			metadata.TagName = BuildTagName(metadata.PathName, metadata.Version)
			metadata.CommitHash = HeadCommit.Hash.String()
			metadata.CreatedAt = HeadCommit.Author.When

			metadataFile.Close()
			ans = append(ans, metadata)
		}
	}

	// 遍历 []Metadata，如果此版本没有在 tag 出现过，则打上新 tag 并 release，同时更新索引文件
	for _, metadata := range ans {
		if exist, err := AddTag(Repo, metadata.TagName); exist {
			oldTag, _ := Repo.Tag(metadata.TagName)
			fmt.Println("Tag already exists: ", metadata.TagName, oldTag.Hash().String())
			continue
		} else if err != nil {
			panic(err)
		}

		//CreateRelease(metadata.TagName)
		NewReleases = append(NewReleases, metadata)
		// 更新索引文件
		NewVersions = append(NewVersions, metadata)
		//AddVersionToIndex(metadata)
	}

	RunCommand("git checkout cache")
	SaveIndex()

	Commit()
	for _, release := range NewReleases {
		CreateRelease(release)
	}
}
