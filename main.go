package main

import (
	"net/http"
	"strings"

	"hobby.com/pkg/repository"
	"hobby.com/pkg/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type category struct {
	Name interface{}
	Page interface{}
}

var sessionCategory category
var allVideo []repository.VideoItem

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("Session", store))

	menu := r.Group("/")
	menu.Use(sessionCheck())

	r.GET("/NSA/video/:name", search)
	r.GET("/NSA/video", search)

	menu.GET("/NSA/moreVideo", getMoreVideo)

	r.Run(":8080")
}

func search(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "夢乃あいか"
	}
	videos, page := service.FindVideos(name, "0", nil)
	collections := service.FindCollections()
	for index := range videos {
		videos[index].PreviewURL = strings.Replace(videos[index].PreviewURL, "https", "http", 1)
	}

	allVideo = videos
	session := sessions.Default(c)
	session.Set("Name", name)
	session.Set("Page", page)
	session.Save()

	c.HTML(200, "videos.html", gin.H{
		"videos":      allVideo,
		"collections": collections,
	})
}

func getMoreVideo(c *gin.Context) {
	name, _ := c.Get("Name")
	page, _ := c.Get("Page")

	videos, newPage := service.FindVideos(name.(string), page.(string), allVideo)
	collections := service.FindCollections()
	for index := range videos {
		videos[index].PreviewURL = strings.Replace(videos[index].PreviewURL, "https", "http", 1)
	}

	allVideo = videos
	session := sessions.Default(c)
	session.Set("Page", newPage)
	session.Save()

	c.HTML(200, "videos.html", gin.H{
		"videos":      allVideo,
		"collections": collections,
	})
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		sessionCategory.Name = session.Get("Name")
		sessionCategory.Page = session.Get("Page")

		if sessionCategory.Name == nil {
			c.Redirect(http.StatusMovedPermanently, "/NSA/video")
			c.Abort()
		} else {
			c.Set("Name", sessionCategory.Name)
			c.Set("Page", sessionCategory.Page)
			c.Next()
		}
	}
}
