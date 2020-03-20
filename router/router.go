package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golango.cn/cc-gin/handler"
	"golango.cn/cc-gin/router/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middleware.Logger())

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		_ = fmt.Errorf("JWT Error", err.Error())
	}

	// 登录
	r.POST("/login", authMiddleware.LoginHandler)

	group := r.Group("/api")
	group.Use(authMiddleware.MiddlewareFunc())
	{
		// 登录用户信息
		group.GET("/user", handler.Get)
	}

	return r
}
