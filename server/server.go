package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

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
	switch req.URL.String() {
	case url_new:
		com, err := parseComment(req)
		if err != nil {
			// TODO return error message
			log.Println("Parsing comment:", err)
		} else {
			s.New(com)
			log.Println(s.tree)
			// TODO return success message
		}
	case url_list:
		key := req.URL.Query()["comment_id"]
		log.Println("Key:", key)
	}
}

func (s *CommentServer) New(c comment.Comment) {
	newId := 0
	for {
		if _, ok := s.tree[newId]; !ok {
			break
		}
		newId = rand.Int()
	}

	c.Id = newId
	c.Score = 1

	s.tree[newId] = &c
	if _, ok := s.tree[c.Parent]; ok {
		s.tree[c.Parent].AddChild(&c)
	}
}
