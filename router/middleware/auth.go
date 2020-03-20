package middleware

import (
	"github.com/gin-gonic/gin"
	"golango.cn/gin-jwt/handler"
	"golango.cn/gin-jwt/pkg/jwtauth"
	"net/http"
	"time"
)

func AuthInit() (*jwtauth.GinJWTMiddleware, error) {

	return &jwtauth.GinJWTMiddleware{
		Key:              []byte("123"),
		Authenticator:    handler.Authenticator,
		Unauthorized:     handler.Unauthorized,
		SigningAlgorithm: "HS256",
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
	}, nil

}
