package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AvGoleCollections struct {
	Success  bool `json:"success"`
	Response struct {
		HasMore          bool             `json:"has_more"`
		TotalCollections int              `json:"total_collections"`
		CurrentOffset    int              `json:"current_offset"`
		Limit            int              `json:"limit"`
		Collections      []CollectionItem `json:"collections"`
	} `json:"response"`
}

type CollectionItem struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Keyword       string `json:"keyword"`
	CoverURL      string `json:"cover_url"`
	TotalViews    int    `json:"total_views"`
	VideoCount    int    `json:"video_count"`
	CollectionURL string `json:"collection_url"`
}

type Video struct {
	Success  bool `json:"success"`
	Response struct {
		HasMore          bool        `json:"has_more"`
		TotalCollections int         `json:"total_collections"`
		CurrentOffset    int         `json:"current_offset"`
		Limit            int         `json:"limit"`
		Videos           []VideoItem `json:"videos"`
	} `json:"response"`
}

type VideoItem struct {
	VID         string `json:"vid"`
	UID         string `json:"uid"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Channel     string `json:"channel"`
	Duration    int    `json:"duration"`
	Framerate   int    `json:"framerate"`
	HD          bool   `json:"hd"`
	AddTime     int    `json:"addtime"`
	ViewNumber  int    `json:"viewnumber"`
	Likes       int    `json:"likes"`
	DisLikes    int    `json:"dislikes"`
	VideoURL    string `json:"video_url"`
	EmbeddedURL string `json:"embedded_url"`
	PreviewURL  string `json:"preview_url"`
}

// AvgleからVideo Collectionを取得
func GetCollections(page string) *AvGoleCollections {
	values := url.Values{}
	values.Add("limit", "20")
	resp, err := http.Get("https://api.avgle.com/v1/collections/" + page + "?" + values.Encode())

	if err != nil {
		fmt.Printf("コレクションエラー%s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var collections AvGoleCollections
	err = json.Unmarshal(body, &collections)
	if err != nil {
		fmt.Println(err)
	}

	return &collections
}

// AvgleからVideo を取得
func GetVideos(word, page string) *Video {
	values := url.Values{}
	values.Add("limit", "250")
	values.Add("o", "mv")
	resp, err := http.Get("https://api.avgle.com/v1/jav/" + word + "/" + page + "?" + values.Encode())

	if err != nil {
		fmt.Printf("コレクションエラー%s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var videos Video
	err = json.Unmarshal(body, &videos)
	if err != nil {
		fmt.Println(err)
	}

	return &videos
}
