package domain

type Request struct {
	Name   string `json:"name"`
	Titles []struct {
		Title string `json:"title"`
	} `json:"titles"`
	Page int `json:"page"`
}

type Responce struct {
	Name   string      `json:"name"`
	Videos []VideoItem `json:"videos"`
	Page   int         `json:"page"`
}

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
	VID         string  `json:"vid"`
	UID         string  `json:"uid"`
	Title       string  `json:"title"`
	Keyword     string  `json:"keyword"`
	Channel     string  `json:"channel"`
	Duration    float64 `json:"duration"`
	Framerate   float64 `json:"framerate"`
	HD          bool    `json:"hd"`
	AddTime     int     `json:"addtime"`
	ViewNumber  int     `json:"viewnumber"`
	Likes       int     `json:"likes"`
	DisLikes    int     `json:"dislikes"`
	VideoURL    string  `json:"video_url"`
	EmbeddedURL string  `json:"embedded_url"`
	PreviewURL  string  `json:"preview_url"`
}
