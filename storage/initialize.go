package storage

import (
	"xiaowei_list/config"
	"xiaowei_list/global"
)

// Setup 配置storage组件
func Setup() {

	// 读取yml配置文件
	global.VP = config.InitViper()

	// 初始化数据库
	global.DB = config.InitMySQL()

}
