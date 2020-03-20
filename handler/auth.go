package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticator(c *gin.Context) (map[string]interface{}, error) {

	var user map[string]interface{}
	c.ShouldBind(&user)

	// 用户名密码
	userName := user["username"].(string)
	userPwd := user["password"].(string)

	// 模拟登录
	if userName == "admin" && userPwd == "123" {
		return map[string]interface{}{
			"name":   "admin",
			"mobile": "15211110000",
		}, nil
	}

	return nil, errors.New("用户名或密码错误")
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
	c.Abort()
}

