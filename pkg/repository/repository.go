package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	. "hobby.com/pkg/domain"
)

// AvgleからVideo Collectionを取得
// func GetCollections() *AvGoleCollections {
// 	values := url.Values{}
// 	values.Add("limit", "70")
// 	values.Add("o", "mv")
// 	resp, err := http.Get("https://api.avgle.com/v1/collections/0?" + values.Encode())

// 	if err != nil {
// 		fmt.Printf("コレクションエラー%s", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}

// 	var collections AvGoleCollections
// 	err = json.Unmarshal(body, &collections)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return &collections
// }

// AvgleからVideo を取得
func GetVideos(word, page string) *Video {
	values := url.Values{}
	values.Add("limit", "40")
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
