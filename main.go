package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jovechiang/lettr/handlers"
)

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
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		phoneNum := c.Query("phone_number")
		pwd := c.Query("pwd")
		resp := handlers.ProcessLogin(phoneNum, pwd)
		resp.LoggedUser.PhoneNum = phoneNum
		jsonResp, _ := json.Marshal(resp)
		c.JSON(200, gin.H{
			"response": string(jsonResp),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
