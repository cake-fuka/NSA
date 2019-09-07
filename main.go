package main

import (
	"strings"

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
	videos := service.FindVideos(name)
	collections := service.FindCollections()
	for index := range videos {
		videos[index].PreviewURL = strings.Replace(videos[index].PreviewURL, "https", "http", 1)
	}
	c.HTML(200, "videos.html", gin.H{
		"videos":      videos,
		"collections": collections,
	})
}
