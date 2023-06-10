package main

import (
	"envoy-golang-filter-hub/config"
	"envoy-golang-filter-hub/internal/global/logx"
	"envoy-golang-filter-hub/internal/global/oauth"
	"envoy-golang-filter-hub/internal/middleware"
	"envoy-golang-filter-hub/internal/module/user"
	"envoy-golang-filter-hub/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		config.Init()
		oauth.Init()
		logs.Init()
		//database.Init()
		//cache.Init()
		//mq.Init()
		//cron.Init()
	})
}

//TODO: Pflag + Viper + Cobra

func main() {
	r := InitRouters()
	utils.PanicIfErr(
		r.Run(config.Get().Host + ":" + config.Get().Port),
	)
}

func InitRouters() *gin.Engine {
	gin.SetMode(string(config.Get().RunMode))

	r := gin.Default()
	r.Use(middleware.LogMiddleware)
	basic := r.Group("/" + config.Get().Prefix)
	user.InitUserRouter(basic)

	return r
}
