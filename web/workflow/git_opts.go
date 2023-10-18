package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"io"
)

var (
	Repo       *git.Repository
	Head       *plumbing.Reference
	HeadCommit *object.Commit
)

func Commit() {
	// Run git Command

	cmds := []string{
		"git checkout cache",
		"git add .",
		fmt.Sprintf("git commit -m \"Committing changes made by %s in GitHub Workflow\"", GitHubActor),
		"git push origin cache",
		"git checkout main",
		"git push origin --tags",
	}

	for _, cmd := range cmds {
		RunCommand(cmd)
	}
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

func ReadFile(path string) string {
	tree, err := HeadCommit.Tree()
	if err != nil {
		panic(err)
	}

	file, err := tree.File(path)
	if err != nil {
		fmt.Println("File not found: ", path)
		return ""
	}

	reader, err := file.Reader()
	if err != nil {
		panic(err)
	}

	ans, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return string(ans)
}
