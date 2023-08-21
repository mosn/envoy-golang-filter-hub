package service

import (
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type IPlugin interface {
	Read(DirPath string) error
	Render() ([]byte, error)
}

type Plugin struct {
	Plugins []plugin
	DirPath string
}

type plugin struct {
	name     string
	category string
	versions []version
}

type version struct {
	version    semver.Version
	commitHash string
}

func (p *Plugin) Read() error {
	// 获取 plugin 目录下所有的子目录
	pluginDirs, err := os.ReadDir(p.DirPath)
	if err != nil {
		return fmt.Errorf("error reading plugin directory: %w", err)
	}

	// 遍历每个插件目录，并解析 metadata.yaml 文件
	for _, pluginDir := range pluginDirs {
		if !pluginDir.IsDir() {
			continue
		}

		pluginPath := filepath.Join(p.DirPath, pluginDir.Name())
		metadataPath := filepath.Join(pluginPath, "metadata.yaml")

		// 解析 metadata.yaml 文件
		pluginData, err := p.parsePluginMetadata(metadataPath)
		if err != nil {
			return fmt.Errorf("error parsing metadata for plugin %s: %w", pluginDir.Name(), err)
		}

		// 获取每个插件的版本信息
		versions, err := p.getPluginVersions(pluginPath)
		if err != nil {
			return fmt.Errorf("error getting versions for plugin %s: %w", pluginDir.Name(), err)
		}

		// 将版本信息存入 Plugin 结构体中
		pluginData.versions = versions
		p.Plugins = append(p.Plugins, pluginData)
	}

	return nil
}

func (p *Plugin) parsePluginMetadata(metadataPath string) (plugin, error) {
	// 读取 metadata.yaml 文件内容
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return plugin{}, err
	}

	// 定义一个结构体来存储 metadata.yaml 中的数据
	var metadata struct {
		Name     string `yaml:"Name"`
		Version  string `yaml:"Version"`
		Category string `yaml:"Category"`
	}

	// 解析 yaml 数据到结构体中
	err = yaml.Unmarshal(data, &metadata)
	if err != nil {
		return plugin{}, err
	}

	// 解析版本号为 semver.Version 类型
	v, err := semver.NewVersion(metadata.Version)
	if err != nil {
		return plugin{}, err
	}

	// 创建 plugin 结构体并返回
	return plugin{
		name:     metadata.Name,
		category: metadata.Category,
		versions: []version{{version: *v}},
	}, nil
}

func (p *Plugin) getPluginVersions(pluginPath string) ([]version, error) {
	var versions []version

	// 打开 Git 仓库
	repo, err := git.PlainOpen(p.DirPath)
	if err != nil {
		return nil, fmt.Errorf("error opening git repository: %w", err)
	}

	// 获取插件目录下的历史提交
	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("error getting head reference: %w", err)
	}
	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, fmt.Errorf("error iterating commits: %w", err)
	}

	// 遍历历史提交，查找每个 commit 中 plugin 目录下的文件内容
	err = commitIter.ForEach(func(commit *object.Commit) error {
		// 获取某个 commit 中 plugin 目录下的文件内容
		tree, err := commit.Tree()
		if err != nil {
			return err
		}
		file, err := tree.File(pluginPath + "/metadata.yaml")
		if err != nil {
			// 如果文件不存在，则跳过该 commit
			if err == object.ErrFileNotFound {
				return nil
			}
			return err
		}

		// 解析 metadata.yaml 文件内容
		data, err := file.Contents()
		if err != nil {
			return err
		}
		var metadata struct {
			Version string `yaml:"Version"`
		}
		err = yaml.Unmarshal([]byte(data), &metadata)
		if err != nil {
			return err
		}

		// 解析版本号为 semver.Version 类型
		v, err := semver.NewVersion(metadata.Version)
		if err != nil {
			return err
		}

		// 将版本信息存入结果数组
		versions = append(versions, version{
			version:    *v,
			commitHash: commit.Hash.String(),
		})

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error iterating commits: %w", err)
	}

	return versions, nil
}

func (p *Plugin) Render() ([]byte, error) {

	//TODO implement me
	panic("implement me")

}
