package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golango.cn/gin-jwt/handler"
	"golango.cn/gin-jwt/router/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middleware.Logger())

	//jwt中间件
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
