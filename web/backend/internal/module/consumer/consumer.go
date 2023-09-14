package consumer

import (
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/global/mq"
	"envoy-go-fliter-hub/internal/module/parse"
	"envoy-go-fliter-hub/internal/module/render"
	"envoy-go-fliter-hub/internal/module/update"
)

func Init() {
	logger := logs.NameSpace("Consumer").Sugar()
	go func() {
		for {
			mq.MQ.ListenUpdateSignal()

			// 更新本地仓库
			repo, err := update.Update.Update()
			if err != nil {
				logger.Error("update error", err)
				continue
			}

			// 重新解析
			data, err := parse.Parse.Parse(repo)
			if err != nil {
				logger.Error("parse error", err)
				continue
			}

			// 重新渲染
			err = render.Render.Render(data)
			if err != nil {
				logger.Error("render error", err)
				continue
			}
		}
	}()

}
