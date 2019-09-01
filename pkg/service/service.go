package service

import (
	"strings"

	"hobby.com/pkg/repository"

	"github.com/bluele/mecab-golang"
)

func FindVideos(word, page string) []repository.VideoItem {
	video := repository.GetVideos(word, page)
	videos := video.Response.Videos
	okVideo := []repository.VideoItem{}
	wordsList := [][]string{}
	for _, v := range videos {
		if !matching(wordsList, v.Title) {
			okVideo = append(okVideo, v)
			wordsList = append(wordsList, analysis(v.Title))
		}
	}
	// jsonModel, _ := json.Marshal(okVideo)
	return okVideo
}

func FindCollections(page string) []repository.CollectionItem {
	collection := repository.GetCollections(page)
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
