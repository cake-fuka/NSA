package main

import (
	"hobby.com/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/collection", func(c *gin.Context) {
		output := service.FindCollections("0")
		c.JSON(200, output)
	})

	r.GET("/video", func(c *gin.Context) {
		ouput := service.FindVideos("夢乃あいか", "0")
		c.JSON(200, ouput)
	})

	r.Run()
}
