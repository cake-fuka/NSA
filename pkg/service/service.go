package service

import (
	"encoding/json"

	"hobby.com/pkg/repository"
)

func FindVideos(word, page string) string {
	video := repository.GetVideos(word, page)
	videos := video.Response.Videos
	jsonModel, _ := json.Marshal(videos)
	return string(jsonModel)
}

func FindCollections(page string) string {
	collection := repository.GetCollections(page)
	collections := collection.Response.Collections
	jsonModel, _ := json.Marshal(collections)
	return string(jsonModel)
}

func analysis() string {
	return ""
}
