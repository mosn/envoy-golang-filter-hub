package update

import (
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	config := Config{
		RemoteRepoUrl: "https://github.com/NX-Official/envoy-golang-filter-hub", // 替换为一个实际的测试用仓库URL
		LocalRepoPath: "./temp_test_repo",                                       // 临时本地路径
		//Username:     "", // 如果需要的话
		//Password:     "", // 如果需要的话
	}

	u := newUpdate(config)

	// 测试克隆和更新仓库
	repo, err := u.Update()
	if err != nil {
		t.Errorf("Failed to clone or update repository: %v", err)
	}

	// 检查是否确实返回了一个仓库对象
	if repo == nil {
		t.Errorf("Expected a repository object, got nil")
	}

	// 检查仓库是否存在于预期的本地路径
	if _, err := os.Stat(config.LocalRepoPath + "/.git"); os.IsNotExist(err) {
		t.Errorf("Repository not found at expected path: %v", config.LocalRepoPath)
	}

	// 移除测试仓库（可选）
	os.RemoveAll(config.LocalRepoPath)
}
