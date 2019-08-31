package service

import (
	"encoding/json"
	"strings"

	"hobby.com/pkg/repository"

	"github.com/bluele/mecab-golang"
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

func analysis(text string) []string {
	m, _ := mecab.New()
	defer m.Destroy()
	tg, _ := m.NewTagger()
	lt, _ := m.NewLattice(text)
	node := tg.ParseToNode(lt)
	words := []string{}
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			words = append(words, node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}
	return words
}
