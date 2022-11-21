package main

import (
	"WebFilesBrowser/app"
	"WebFilesBrowser/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	_ = server.SetTrustedProxies(nil)
	server.Use(handle) // 异常处理
	server.StaticFS("/static", http.Dir("./static"))
	server.StaticFS("/files", http.Dir(config.App.Path))
	apiRoute(server.RouterGroup.Group("/api"))
	err := server.Run(config.App.Host + ":" + config.App.Port) // 启动服务
	if err != nil {
		fmt.Println(err)
	}
}

func apiRoute(api *gin.RouterGroup) {
	files := new(app.Files)

	api.GET("/get/files/list", files.GetFilesList)
}

// handle 异常处理
func handle(c *gin.Context) {
	c.Next()
	e := recover() // 捕获异常
	if e != nil {
		c.JSON(500, gin.H{"code": 500, "msg": toString(e), "data": nil})
		c.Abort()
	}
}

// toString error 转 string
func toString(e interface{}) string {
	switch v := e.(type) {
	case error:
		return v.Error()
	default:
		return e.(string)
	}
}
