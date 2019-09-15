package main

import (
	"net/http"

	. "hobby.com/pkg/domain"
	"hobby.com/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.POST("/NSA/video", search)

	r.Run(":8080")
}

func search(c *gin.Context) {
	var req Request
	c.BindJSON(&req)
	if req.Name == "" {
		req.Name = "夢乃あいか"
	}
	resp := service.FindVideos(req)
	resp.Name = req.Name
	c.JSON(http.StatusOK, resp)
}
