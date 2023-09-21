package update

import (
	"github.com/go-git/go-git/v5"
	"os"
)

func (u update) Update() (*git.Repository, error) {
	remoteRepoUrl := u.config.RemoteRepoUrl
	localRepoPath := u.config.LocalRepoPath

	// 尝试打开本地仓库
	repo, err := git.PlainOpen(localRepoPath)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			// 如果本地仓库不存在，则克隆远程仓库
			repo, err = git.PlainClone(localRepoPath, false, &git.CloneOptions{
				URL: remoteRepoUrl,
				//Auth: &http.BasicAuth{
				//	Username: u.config.Username, // 这里可以是任何内容，因为GitHub不使用这个字段
				//	Password: u.config.Password,
				//},
				Progress: os.Stdout,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// 如果本地仓库存在，则拉取更新
		workTree, err := repo.Worktree()
		if err != nil {
			return nil, err
		}

		err = workTree.Pull(&git.PullOptions{
			RemoteName: "origin",
			//Auth: &http.BasicAuth{
			//	Username: u.config.Username,
			//	Password: u.config.Password,
			//},
			Progress: os.Stdout,
		})
		// 当仓库已经是最新状态时，Pull操作会返回"已经是最新"的错误，这里我们忽略它
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return nil, err
		}
	}
	return repo, nil
}
