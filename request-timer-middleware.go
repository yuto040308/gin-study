package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestTimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 前処理

		start := time.Now()

		c.Next() // 次のミドルウェアまたはハンドラへ処理が移る

		// 後処理
		duration := time.Since(start)
		log.Printf("Request took %s", duration)
	}
}
