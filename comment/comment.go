package comment

type Comment struct {
	Text     string `json:"text"`
	Parent   int    `json:"parent"`
	Id       int    `json:"id"`
	Children []int  `json:"children"`
	Score    int    `json:"score"`
}
