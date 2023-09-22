package main

import (
	"context"
	"encoding/json"
	"envoy-go-filter-hub/model"
	"envoy-go-filter-hub/template"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v55/github"
	"github.com/mholt/archiver/v3"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	GitHubToken   = os.Getenv("GITHUB_TOKEN")
	GitHubRepo    = os.Getenv("GITHUB_REPOSITORY")
	GitHubRepoUrl = "https://github.com/" + GitHubRepo
	GitHubActor   = os.Getenv("GITHUB_ACTOR")
	GitHubClient  = github.NewClient(nil).WithAuthToken(GitHubToken)
	Repo          *git.Repository
	Head          *plumbing.Reference
	HeadCommit    *object.Commit
	RootPath      = filepath.Join("../../")
	IndexPath     = filepath.Join(RootPath, "web/cache/index.json")
	PluginMap     = make(map[string]template.PluginBasic)
	NewReleases   = make([]model.Metadata, 0)
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
		AddVersionToIndex(metadata)
	}
	SaveIndex()

	Commit()
	for _, release := range NewReleases {
		CreateRelease(release)
	}
}

func Commit() {
	// Run git Command

	cmds := []string{
		"git add .",
		fmt.Sprintf("git commit -m \"Committing changes made by %s in GitHub Workflow\"", GitHubActor),
		"git push origin main --tags",
	}

	for _, cmd := range cmds {
		cmd := exec.Command("bash", "-c", cmd)
		cmd.Dir = RootPath
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		//if err := cmd.Run(); err != nil {
		//	panic(err) // When nothing to commit, it will panic
		//}
	}
}

func BuildTagName(pluginName string, version string) string {
	return fmt.Sprintf("%s|v%s", pluginName, version)
}

func AddTag(r *git.Repository, tagName string) (bool, error) {
	_, err := r.CreateTag(tagName, Head.Hash(), nil)
	if err == git.ErrTagExists {
		return true, nil
	} else if err != nil {
		return false, err
	} else {
		return false, nil
	}
}

func CreateRelease(r model.Metadata) {
	fmt.Println("Creating release for tag: ", r.TagName)
	// 准备 release 信息
	releaseTitle := fmt.Sprintf("%s - v%s", r.Name, r.Version) // 替换为你的 release 标题
	releaseBody := ""

	// 创建 release
	release := &github.RepositoryRelease{
		TagName: &r.TagName,
		Name:    &releaseTitle,
		Body:    &releaseBody,
	}

	owner, repoName, found := strings.Cut(GitHubRepo, "/")
	if !found {
		panic("Error: Not Found GITHUB_REPOSITORY")
	}

	createdRelease, _, err := GitHubClient.Repositories.CreateRelease(context.Background(), owner, repoName, release)
	if err != nil {
		log.Fatal(err)
	}

	// 打包目录
	pluginDir := filepath.Join(RootPath, "plugins", r.PathName)
	//pluginPathName := "."
	zipFileName := fmt.Sprintf("%s v%s", r.Name, r.Version) + ".zip"

	if err := archiver.Archive([]string{pluginDir}, zipFileName); err != nil {
		panic(err)
	}

	// 上传压缩文件作为附件
	attachmentFile, err := os.Open(zipFileName)
	if err != nil {
		panic(err)
	}
	defer attachmentFile.Close()

	attachmentName := zipFileName // 替换为你想要的附件名称

	opt := &github.UploadOptions{
		Name: attachmentName,
	}

	_, _, err = GitHubClient.Repositories.UploadReleaseAsset(context.Background(), owner, repoName, *createdRelease.ID, opt, attachmentFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Uploaded attachment: %s\n", attachmentName)

	// 清理临时压缩文件
	if err := os.Remove(zipFileName); err != nil {
		log.Println("Error deleting zip file:", err)
	}

	fmt.Printf("Created release: %s\n", *createdRelease.Name)
}

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

	pluginDetailPath := filepath.Join(RootPath, "web/cache/plugins", fmt.Sprintf("%s.json", metadata.PathName))
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

func RenderMarkdown(markdown string) string {
	// 使用 GitHub API https://api.github.com/markdown
	// 请见 https://docs.github.com/zh/free-pro-team@latest/rest/markdown/markdown?apiVersion=2022-11-28#render-a-markdown-document

	renderedHTML, _, err := GitHubClient.Markdown(context.Background(), markdown,
		&github.MarkdownOptions{
			Mode:    "gfm",
			Context: GitHubRepo,
		})
	if err != nil {
		panic(err)
	}
	return renderedHTML
}

func SaveIndex() {
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
