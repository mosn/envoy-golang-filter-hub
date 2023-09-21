package consumer

import (
	"envoy-go-fliter-hub/internal/module/parse"
	"envoy-go-fliter-hub/internal/module/render"
	"envoy-go-fliter-hub/internal/module/update"
	"fmt"
)

func updateConsumer() error {
	// 更新本地仓库
	repo, err := update.Update.Update()
	if err != nil {
		return err
	}

	fmt.Printf("repo: %+v\n", repo)

	// 重新解析
	data, err := parse.Parse.Parse(repo)
	if err != nil {
		return err
	}

	fmt.Printf("data: %+v\n", data)

	// 重新渲染
	err = render.Render.Render(data)
	if err != nil {
		return err
	}

	return nil
}
