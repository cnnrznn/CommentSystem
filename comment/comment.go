package comment

import (
	"fmt"
)

type Comment struct {
	Text     string     `json:"text"`
	Parent   int        `json:"parent"`
	Id       int        `json:"id"`
	Children []*Comment `json:"children"`
	Score    int        `json:"score"`
}

func (c *Comment) AddChild(child *Comment) {
	// TODO lock mutex
	c.Children = append(c.Children, child)
	// TODO unlock mutex
}

func (c *Comment) String() string {
	return fmt.Sprintf("{%v: \n\t%v\n\t\n}", c.Text, c.Children)
}
