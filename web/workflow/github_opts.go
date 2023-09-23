package main

import (
	"context"
	"envoy-go-filter-hub/model"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/mholt/archiver/v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	GitHubToken   = os.Getenv("GITHUB_TOKEN")
	GitHubRepo    = os.Getenv("GITHUB_REPOSITORY")
	GitHubRepoUrl = "https://github.com/" + GitHubRepo
	GitHubActor   = os.Getenv("GITHUB_ACTOR")
	GitHubClient  = github.NewClient(nil).WithAuthToken(GitHubToken)
)

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
