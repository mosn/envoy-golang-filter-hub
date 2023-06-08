package main

import (
	"envoy-golang-filter-hub/config"
	"envoy-golang-filter-hub/internal/global/log"
	"envoy-golang-filter-hub/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		config.Init()
		log.Init()
		//database.Init()
		//cache.Init()
		//mq.Init()
		//cron.Init()
	})
}

func main() {
	r := gin.Default()

	gin.SetMode(string(config.Get().RunMode))

	utils.PanicIfErr(
		r.Run(config.Get().Host + ":" + config.Get().Port),
	)

}
