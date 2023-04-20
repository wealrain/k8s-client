package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// 不拦截登录接口
	if c.Request.URL.Path == "/token/create" {
		c.Next()
		return
	}

	// 从请求头中获取token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token为空"})
		return
	}

	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("ljyun"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	// 验证token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 判断token是否过期
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token已过期"})
			return
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
	}
}
