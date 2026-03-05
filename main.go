package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello gin"})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"users": []string{"taro", "hanako"}})
	})

	// パスパラメーター受け取り
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"id": id})
	})

	// クエリパラメーター（?q=foo）受け取り
	r.GET("search", func(c *gin.Context) {
		q := c.Query("q") // 例: /search?q=gin
		c.JSON(200, gin.H{"query": q})
	})

	// リクエストボディ（JSON）
	// JSONを送ると構造体にバインドされる
	type CreateUserInput struct {
		Name string `json:"name"`
	}
	r.POST("/users", func(c *gin.Context) {
		var input CreateUserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "created", "name": input.Name})
	})

	/*
		r.POST("/users", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "user created"})
		})*/

	// デフォルトで8080
	r.Run()
}
