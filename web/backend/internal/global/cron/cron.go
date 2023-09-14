package cron

import (
	"envoy-go-fliter-hub/internal/global/mq"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Init() {
	go func() {
		c := cron.New()

		// 每天凌晨 0 点执行一次
		_, err := c.AddFunc("@every 1day", mq.MQ.SendUpdateSignal)

		if err != nil {
			fmt.Println("添加定时任务失败:", err)
			panic(err)
			return
		}

		// 启动 cron 调度器
		c.Start()
	}()
}
