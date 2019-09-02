package service

import (
	"strconv"
	"strings"

	"hobby.com/pkg/repository"

	"github.com/bluele/mecab-golang"
)

func FindVideos(word string) []repository.VideoItem {
	okVideo := []repository.VideoItem{}
	page := 0
	for len(okVideo) < 30 {
		video := repository.GetVideos(word, strconv.Itoa(page))
		videos := video.Response.Videos
		wordsList := [][]string{}
		for _, v := range videos {
			if !matching(wordsList, v.Title) {
				okVideo = append(okVideo, v)
				wordsList = append(wordsList, analysis(v.Title))
			}
		}
		if len(videos) < 40 {
			break
		}
		page++
	}
	return okVideo
}

func FindCollections() []repository.CollectionItem {
	collection := repository.GetCollections()
	collections := collection.Response.Collections
	// jsonModel, _ := json.Marshal(collections)
	return collections
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

func matching(videos [][]string, title string) bool {
	max := 0.0
	for _, x := range videos {
		cnt := 0.0
		length := len(x)
		for _, y := range x {
			if strings.Contains(title, y) {
				cnt += 1 / float64(length)
			}
		}
		if max < cnt {
			max = cnt
		}
	}
	if max >= 0.5 {
		return true
	} else {
		return false
	}
}
