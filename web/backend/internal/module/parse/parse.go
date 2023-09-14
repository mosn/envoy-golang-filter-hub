package parse

import (
	"envoy-go-fliter-hub/internal/model"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
)

// Parse 函数用于解析 git 标签和关联的元数据
func (p parse) Parse(r *git.Repository) ([]model.Metadata, error) {
	var ans []model.Metadata // 用于存储所有解析出的元数据

	// 从仓库中获取所有标签
	tags, err := r.Tags()
	if err != nil {
		return nil, errors.Wrap(err, "无法获取标签")
	}

	// 遍历每一个标签
	err = tags.ForEach(func(t *plumbing.Reference) error {
		tagName := t.Name().Short()
		pathName, err := getPathName(tagName)
		if err != nil {
			return errors.Wrapf(err, "无法解析标签 %s", tagName)
		}

		// 获取该标签对应的提交对象
		commit, err := r.CommitObject(t.Hash())
		if err != nil {
			return errors.Wrap(err, "获取提交对象失败")
		}

		// 获取该提交的树对象
		tree, err := commit.Tree()
		if err != nil {
			return errors.Wrap(err, "获取树对象失败")
		}

		// 在树对象中查找 plugin 目录
		pluginTree, err := tree.Tree("plugins")
		if err != nil {
			return errors.Wrap(err, "找不到 plugin 目录")
		}

		// 获取 pathName/metadata.yaml 文件
		metadataFile, err := pluginTree.File(fmt.Sprintf("%s/metadata.yaml", pathName))
		if err != nil {
			return errors.Wrapf(err, "找不到 metadata.yaml 文件在 %s/metadata.yaml", pathName)
		}

		// 读取 metadata.yaml 文件
		reader, err := metadataFile.Reader()
		if err != nil {
			return errors.Wrap(err, "文件读取失败")
		}
		defer reader.Close()

		// 解析 metadata.yaml 文件内容
		metadata, err := parseMetadata(reader)
		if err != nil {
			return err
		}
		metadata.TagName = tagName
		metadata.CommitHash = commit.Hash.String()
		metadata.PathName = pathName
		metadata.CreatedAt = commit.Author.When

		// 将解析出的元数据添加到 ans 切片中
		ans = append(ans, metadata)
		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "无法遍历标签")
	}

	return ans, nil
}

// getPathName 函数用于从标签名中解析出路径名
func getPathName(tag string) (string, error) {
	parts := strings.Split(tag, "|")
	if len(parts) < 2 {
		return "", errors.New("标签格式无效")
	}

	return parts[0], nil
}

// parseMetadata 函数用于解析 metadata.yaml 文件内容
func parseMetadata(reader io.ReadCloser) (model.Metadata, error) {
	var metadata model.Metadata

	content, err := io.ReadAll(reader)
	if err != nil {
		return metadata, errors.Wrap(err, "读取内容失败")
	}

	err = yaml.Unmarshal(content, &metadata)
	if err != nil {
		return metadata, errors.Wrap(err, "内容反序列化失败")
	}

	return metadata, nil
}
