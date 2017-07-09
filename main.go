package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "Hello World! aLettr",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "hit register",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
