package types

// ZhihuTopSearch 即知乎热搜榜
type ZhihuTopSearch struct {
	TopSearch TopSearch `json:"top_search"`
}

// TopSearch 即每个
type TopSearch struct {
	Words []aTopSearch `json:"words"`
}

type aTopSearch struct {
	Query string `json:"query"`
	DisplayQuery string `json:"display_query"`
}

