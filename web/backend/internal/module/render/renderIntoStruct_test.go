package render

import (
	"encoding/json"
	"envoy-go-fliter-hub/internal/model"
	"envoy-go-fliter-hub/internal/template"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRenderIntoStruct(t *testing.T) {
	r := render{}

	// 准备测试数据
	metadata := []model.Metadata{
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
	details, list, err := r.renderIntoStruct(metadata)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// 验证 PluginDetail
	expectedDetail := template.PluginDetail{
		PathName:    "example",
		Name:        "New ExamplePlugin",
		Version:     "1.2",
		Category:    "Utility",
		Description: "This is an example plugin",
		Versions: []template.Version{
			{
				Version:    "1.1",
				CreatedAt:  "2021-01-01",
				CommitHash: "abc123",
			},
			{
				Version:    "1.2",
				CreatedAt:  "2021-02-01",
				CommitHash: "def456",
			},
		},
	}

	//fmt.Printf("%+v\n", details[0])
	jsonData, err := json.MarshalIndent(details, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	if !reflect.DeepEqual(details[0], expectedDetail) {
		t.Errorf("Expected %v, but got %v", expectedDetail, details[0])
	}

	// 验证 PluginList
	expectedList := template.PluginList{
		TotalCount: 1,
		Plugins: []template.PluginBasic{
			{
				PathName:    "example",
				Name:        "New ExamplePlugin",
				Version:     "1.2",
				Category:    "Utility",
				Description: "This is an example plugin",
			},
		},
	}

	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("Expected %v, but got %v", expectedList, list)
	}

}
