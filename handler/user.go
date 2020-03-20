package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// 获取用户信息
// @Summary 获取用户信息
// @Tags 模块
// @Accept  json
// @Produce  json
func Get(c *gin.Context) {

	v, _ := c.Get("JWT_PAYLOAD")
	claims := v.(jwt.MapClaims)

	c.JSON(200, gin.H{
		"message": "ok",
		"user":    claims,
	})

}
