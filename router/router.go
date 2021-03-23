package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r := gin.New()
	gin.ForceConsoleColor()
	//gin.SetMode(gin.ReleaseMode)
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signUp", controller.SignUpHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "404"})
	})
	return r
}
