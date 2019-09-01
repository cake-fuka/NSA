package main

import (
	"hobby.com/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/video/:name", search)
	r.GET("/video", search)

	r.Run(":8080")
}

func search(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "夢乃あいか"
	}
	videos := service.FindVideos(name, "0")
	collections := service.FindCollections("0")
	c.HTML(200, "videos.html", gin.H{
		"videos":      videos,
		"collections": collections,
	})
}
