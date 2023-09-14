package render

import (
	"envoy-go-fliter-hub/internal/module/parse"
	"testing"
	"time"
)

func Test_render_Render(t *testing.T) {
	r := newRender(Config{
		OutPutDir: "/Users/nx/GolandProjects/OSPP2023/envoy-golang-filter-hub/web/backend/static/output",
	})

	// 准备测试数据
	metadata := []parse.Metadata{
		{
			TagName:     "example|v1.1",
			CommitHash:  "abc123",
			CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Name:        "ExamplePlugin",
			PathName:    "example",
			Description: "This is an example plugin",
			Category:    "Utility",
			Version:     "1.1",
		},
		{
			TagName:     "example|v1.2",
			CommitHash:  "def456",
			CreatedAt:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			Name:        "New ExamplePlugin",
			PathName:    "example",
			Description: "This is an example plugin",
			Category:    "Utility",
			Version:     "1.2",
		},
	}

	// 调用待测试的方法
	err := r.Render(metadata)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
