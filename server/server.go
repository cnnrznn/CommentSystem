package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/cnnrznn/comment/comment"
)

const (
	url_new  string = "/Comment/New"
	url_list string = "/Comment/List"
)

func main() {
	cs := NewCommentServer()

	mux := http.NewServeMux()
	mux.Handle(url_new, cs)
	mux.Handle(url_list, cs)

	log.Fatal(http.ListenAndServe(":8888", mux))
}

type CommentServer struct {
	tree map[int]*comment.Comment
}

func NewCommentServer() *CommentServer {
	cs := CommentServer{
		tree: make(map[int]*comment.Comment),
	}
	cs.tree[0] = &comment.Comment{}
	cs.tree[0].Text = "<Post text here>"

	return &cs
}

func parseComment(req *http.Request) (c comment.Comment, err error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &c)

	return
}

func (s *CommentServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case url_new:
		com, err := parseComment(req)
		if err != nil {
			// TODO return error message
			log.Println("Parsing comment:", err)
		} else {
			cId := comment.Id{s.New(com)}
			bytes, err := json.Marshal(cId)
			if err != nil {
				log.Println("Marshal-ing id:", err)
			}
			w.Write(bytes)
		}
	case url_list:
		key, err := strconv.Atoi(req.URL.Query()["comment_id"][0])
		if err != nil {
			// TODO return error
			return
		}

		if obj, ok := s.tree[key]; ok {
			bytes, err := json.Marshal(obj)
			if err != nil {
				log.Println("Marshal-ing comment:", err)
				return
			}
			w.Write(bytes)
		}
	}
}

func (s *CommentServer) New(c comment.Comment) int {
	// TODO lock tree for parallel random number generation

	if _, ok := s.tree[c.Parent]; !ok {
		log.Println("Parent doesn't exist in tree")
		return -1
	}

	newId := 0
	for {
		if _, ok := s.tree[newId]; !ok {
			break
		}
		newId = rand.Int()
	}

	c.Id = newId
	c.Score = 1

	s.tree[c.Parent].AddChild(&c)
	s.tree[newId] = &c

	return newId
}
