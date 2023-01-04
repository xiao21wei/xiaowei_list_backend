package main

import (
	"github.com/gin-gonic/gin"
	"xiaowei_list/global"
	"xiaowei_list/router"
	"xiaowei_list/storage"
)

func main() {
	// 初始化全局资源
	storage.Setup()
	// 创建Router
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":" + global.VP.GetString("server.backend_port"))
	if err != nil {
		return
	}
}
