package cmd

import (
	"envoy-go-fliter-hub/config"
	"envoy-go-fliter-hub/internal/global/cron"
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/global/mq"
	"envoy-go-fliter-hub/internal/module"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once

func Init() {
	once.Do(func() {
		config.Init()
		logs.Init()
		//database.Init()
		//cache.Init()
		mq.Init()
		cron.Init()
		module.Init()
	})
}

func Run() {
	r := gin.Default()
	basic := r.Group("/" + config.Config.Prefix)

	err := r.Run(config.Config.Host + ":" + config.Config.Port)
	if err != nil {
		panic(err)
	}
}
