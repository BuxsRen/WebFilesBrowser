package app

import (
	"WebFilesBrowser/config"
	"WebFilesBrowser/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type Files struct {
}

// GetFilesList 读取文件列表
func (f Files) GetFilesList(c *gin.Context) {

	dir := c.Query("dir")

	dir = strings.Replace(dir, "./", "", -1)
	dir = strings.Replace(dir, "../", "", -1)

	list, err := utils.GetFilesList(config.App.Path + "/" + dir)
	if err != nil {
		c.JSON(500, utils.Error(500, err.Error()))
		return
	}

	c.JSON(200, utils.JSON("获取成功", list))
}
