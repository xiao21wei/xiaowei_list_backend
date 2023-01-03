package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
	v1 "xiaowei_list/api/v1"
	_ "xiaowei_list/docs"
	"xiaowei_list/global"
	"xiaowei_list/middleware"
)

func InitRouter(r *gin.Engine) {
	//跨域配置
	r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/hello", HelloGin)
	// 禁用代理访问
	if err := r.SetTrustedProxies(nil); err != nil {
		//global.LOG.Panic("初始化失败：禁止使用代理访问失败")
	}

	// 登录模块
	rawRouter := r.Group("/api/v1")
	{
		rawRouter.POST("/login", v1.Login)
		rawRouter.POST("/register", v1.Register)
	}
	// 静态资源模块
	resourceRouter := rawRouter.Group("/resource")
	{
		resourceRouter.POST("/upload", v1.UploadFile)
		resourceRouter.Static("/file", filepath.Join(global.VP.GetString("root_path"), "resource"))
	}
	// 除了登录模块和静态资源之外，都需要身份认证
	basicRouter := rawRouter.Group("/")
	basicRouter.Use(middleware.AuthRequired())

	// 用户模块
	userRouter := basicRouter.Group("/users")
	{
		userRouter.GET("/:id", v1.GetUser)
	}
}

func HelloGin(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "HelloGin and HelloAir!",
	})
}
