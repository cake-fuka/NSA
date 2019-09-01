package main

import (
	"hobby.com/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/collection", func(c *gin.Context) {
		output := service.FindCollections("0")
		c.HTML(200, "index.html", gin.H{
			"collection": output,
		})
	})

	r.GET("/video", func(c *gin.Context) {
		output := service.FindVideos("夢乃あいか", "0")
		c.HTML(200, "videos.html", gin.H{
			"videos": output,
		})
	})

	r.Run(":8080")
}
