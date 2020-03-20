package jwtauth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type GinJWTMiddleware struct {

	// 加密Key
	Key []byte

	// 授权
	Authenticator func(c *gin.Context) (map[string]interface{}, error)

	// 未授权
	Unauthorized func(*gin.Context, int, string)

	// 加密方式
	SigningAlgorithm string

	// 响应输出方法
	LoginResponse func(*gin.Context, int, string, time.Time)
}

// 登录
func (mw *GinJWTMiddleware) LoginHandler(c *gin.Context) {

	//登录授权
	data, err := mw.Authenticator(c)

	if err != nil {
		mw.Unauthorized(c, 400, err.Error())
		return
	}

	// 生成Token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range data {
		claims[key] = value
	}

	expire := time.Now().Add(time.Hour * 1)
	claims["exp"] = expire.Unix()

	tokenString, err := token.SignedString(mw.Key)

	if err != nil {
		mw.Unauthorized(c, http.StatusOK, err.Error())
		return
	}

	mw.LoginResponse(c, http.StatusOK, tokenString, expire)
}

// 验证权限中间件
func (mw *GinJWTMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.middlewareImpl(c)
	}
}

// 验证权限
func (mw *GinJWTMiddleware) middlewareImpl(c *gin.Context) {

	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		mw.Unauthorized(c, http.StatusUnauthorized, err.Error())
		return
	}

	if claims["exp"] == nil {
		mw.Unauthorized(c, http.StatusBadRequest, "missing exp field")
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		mw.Unauthorized(c, http.StatusBadRequest, "exp must be float64 format")
		return
	}
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		mw.Unauthorized(c, http.StatusUnauthorized, "token is expired")
		return
	}
	c.Set("JWT_PAYLOAD", claims)
	c.Next()

}

// JWT中获取身份信息
func (mw *GinJWTMiddleware) GetClaimsFromJWT(c *gin.Context) (jwt.MapClaims, error) {

	tokenString, err := mw.jwtFromHeader(c, "Authorization")

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mw.Key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil

	} else {
		return nil, err
	}

}

// 读取Header中的Authorization
func (mw *GinJWTMiddleware) jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", errors.New("auth header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("auth header is invalid")
	}

	return parts[1], nil
}
