package main

import "github.com/gin-gonic/gin"

// 認証ミドルウェアのサンプル
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "secret-token" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			// これ以降の処理は実行されない
			c.Abort()
		}
		c.Next()
	}
}
