package consumer

import (
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/global/mq"
	"envoy-go-fliter-hub/internal/module/parse"
	"envoy-go-fliter-hub/internal/module/render"
	"fmt"
	"github.com/go-git/go-git/v5"
)

func Consume() {
	logger := logs.NameSpace("Consumer").Sugar()
	repo, err := git.PlainOpen("/Users/nx/GolandProjects/OSPP2023/envoy-golang-filter-hub")
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	for {
		mq.MQ.ListenUpdateSignal()

		// 重新解析
		data, err := parse.Parse.Parse(repo)

		// 重新渲染
		err = render.Render.Render(data)
		if err != nil {
			logger.Error("render error", err)
			return
		}
	}
}
