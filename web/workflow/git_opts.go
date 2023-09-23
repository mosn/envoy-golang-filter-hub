package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"os/exec"
)

var (
	Repo       *git.Repository
	Head       *plumbing.Reference
	HeadCommit *object.Commit
)

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
