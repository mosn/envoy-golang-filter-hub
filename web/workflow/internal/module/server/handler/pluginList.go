package handler

import (
	"envoy-go-fliter-hub/config"
	"envoy-go-fliter-hub/internal/global/errs"
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/module/render"
	"envoy-go-fliter-hub/tools"
	"github.com/gin-gonic/gin"
	"path"
)

func (h handler) PluginList(c *gin.Context) {

	// 检查文件是否存在
	indexPath := path.Join(config.Config.Repo.IndexOutPutPath, render.PluginListFileName)
	logs.NameSpace("handler").Sugar().Info("indexPath: ", indexPath)

	// 读取文件
	bytes, err := tools.Read(indexPath)
	if err != nil {
		errs.Fail(c, errs.IndexNotFound)
		logs.NameSpace("handler").Sugar().Error("read index file error: ", err)
		return
	}

	// 返回
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, "{\n    \"code\": 200,\n    \"msg\": \"Success\",\n    \"data\": "+string(bytes)+"}")
	//errs.Success(c, string(bytes))
}
