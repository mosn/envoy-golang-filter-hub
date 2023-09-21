package main

import (
	"archive/zip"
	"context"
	"encoding/json"
	"envoy-go-fliter-hub/model"
	"envoy-go-fliter-hub/template"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v55/github"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	//PluginMap      = make(map[string]template.PluginDetail)
	PluginIndexMap = make(map[string]template.PluginBasic)
	//NewTags        = make([]string, 0)
	NewReleases = make([]model.Metadata, 0)
)

func init() {
	//// 创建 GitHub 客户端
	//GitHubClient = github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: GitHubToken},
	//)))
	//exists := false
	//GitHubToken, exists = os.LookupEnv("GITHUB_TOKEN")
	//if !exists {
	//	fmt.Println("Error: Not Found GITHUB_TOKEN")
	//	os.Exit(1)
	//}

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

	// 读取 plugin_list.json 文件
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
		PluginIndexMap[plugin.PathName] = plugin
	}

	//fmt.Println(PluginIndexMap)
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
	exec.Command("git add .")
	exec.Command(fmt.Sprintf("git commit -m \"Committing changes made by %s in GitHub Workflow\"", GitHubActor))
	exec.Command("git push origin main --tags")
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
	pluginPathName := "."
	zipFileName := fmt.Sprintf("%s v%s", r.Name, r.Version) + ".zip"

	if err := zipDirectory(pluginDir, pluginPathName, zipFileName); err != nil {
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

	PluginIndexMap[metadata.PathName] = newPluginBasic

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

	fmt.Println(string(pluginDetailBytes))

	err = os.WriteFile(pluginDetailPath, pluginDetailBytes, 0644)
	if err != nil {
		panic(err)
	}

}

func RenderMarkdown(markdown string) string {
	// 使用 GitHub API https://api.github.com/markdown
	// 请见 https://docs.github.com/zh/free-pro-team@latest/rest/markdown/markdown?apiVersion=2022-11-28#render-a-markdown-document

	//reqUrl := url.URL{
	//	Scheme: "https",
	//	Host:   "api.github.com",
	//	Path:   "markdown",
	//}
	//
	//reqDto := struct {
	//	Text    string `json:"text"`
	//	Mode    string `json:"mode"`
	//	Context string `json:"context"`
	//}{
	//	Text:    markdown,
	//	Mode:    "gfm",
	//	Context: GitHubRepo,
	//}
	//
	//reqBytes, err := json.Marshal(reqDto)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, body, errs := gorequest.New().
	//	Post(reqUrl.String()).
	//	AppendHeader("Authorization", fmt.Sprintf("Bearer %s", GitHubToken)).
	//	Send(string(reqBytes)).
	//	Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized).
	//	End()
	//if errs != nil {
	//	panic(errs)
	//}
	//
	//return body

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

	pluginList.Plugins = make([]template.PluginBasic, 0, len(PluginIndexMap))
	for _, plugin := range PluginIndexMap {
		pluginList.Plugins = append(pluginList.Plugins, plugin)
	}
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

func zipDirectory(sourceDir, directoryToZip, zipFileName string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历目录并将文件添加到ZIP
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		if relPath == directoryToZip {
			// 添加目标目录下的文件到ZIP
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			zipFile, err := archive.Create(relPath)
			if err != nil {
				return err
			}

			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
