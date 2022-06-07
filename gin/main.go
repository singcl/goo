package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(http.StatusOK, "https.tmpl", gin.H{
			"status": "success",
		})
	})

	// https://gin-gonic.com/zh-cn/docs/examples/ascii-json/
	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<button/>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// 监听并在 0.0.0.0:9000 上启动服务
	r.Run(":9000")
	// 监听并在 https://127.0.0.1:8080 上启动服务
	// r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
