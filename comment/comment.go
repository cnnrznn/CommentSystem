package comment

import (
	"fmt"
)

type Comment struct {
	Text     string `json:"text"`
	Parent   int    `json:"parent"`
	Id       int    `json:"id"`
	Children []int  `json:"children"`
	Score    int    `json:"score"`
}

func (c *Comment) AddChild(child int) {
	c.Children = append(c.Children, child)
}

func (c *Comment) String() string {
	return fmt.Sprintf("%v:'%v'", c.Parent, c.Text)
}
