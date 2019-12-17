package comment

import (
	"fmt"
	"sync"
)

type Comment struct {
	Text     string     `json:"text"`
	Parent   int        `json:"parent"`
	Id       int        `json:"id"`
	Children []*Comment `json:"children"`
	Score    int        `json:"score"`
	mux      sync.Mutex
}

func (c *Comment) AddChild(child *Comment) {
	c.mux.Lock()
	c.Children = append(c.Children, child)
	c.mux.Unlock()
}

func (c *Comment) String() string {
	return fmt.Sprintf("{%v: \n\t%v\n\t\n}", c.Text, c.Children)
}
