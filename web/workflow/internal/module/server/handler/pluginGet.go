package handler

import (
	"envoy-go-fliter-hub/config"
	"envoy-go-fliter-hub/internal/global/errs"
	"envoy-go-fliter-hub/internal/module/render"
	"envoy-go-fliter-hub/tools"
	"github.com/gin-gonic/gin"
	"path"
)

func (h handler) PluginGet(c *gin.Context) {
	pathName := c.Param("id")

	// 检查文件是否存在
	pluginPath := path.Join(config.Config.Repo.IndexOutPutPath, render.PluginDetailDir, pathName+render.PluginDetailSuffix)
	//logs.NameSpace("handler").Sugar().Info("indexPath: ", indexPath)

	// 读取文件
	bytes, err := tools.Read(pluginPath)
	if err != nil {
		errs.Fail(c, errs.PluginNotFound.WithTips(pathName))
		//logs.NameSpace("handler").Sugar().Error("read index file error: ", err)
		return
	}

	// 返回
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, "{\n    \"code\": 200,\n    \"msg\": \"Success\",\n    \"data\": "+string(bytes)+"}")
}
