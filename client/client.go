package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/cnnrznn/comment/comment"
)

const (
	url_comment      string = "http://localhost:8888/Comment"
	url_comment_new  string = url_comment + "/New"
	url_comment_list string = url_comment + "/List?comment_id=0"
)

func reqEndpoint(endpoint, method string, data []byte) (body []byte, err error) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Request creation:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP response status =", resp.StatusCode)

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	return
}

func main() {
	var cms comment.Comment
	var newId comment.Id

	c := comment.Comment{
		Text:   "Wow what a cool article!",
		Parent: 0,
	}

	ids := []int{0}

	for i := 0; i < 50; i++ {
		bytes, _ := json.Marshal(c)

		bytes, _ = reqEndpoint(url_comment_new, "POST", bytes)
		json.Unmarshal(bytes, &newId)
		// add id to pool of available
		ids = append(ids, newId.Id)

		c.Parent = ids[rand.Intn(len(ids))]

		bytes, _ = reqEndpoint(url_comment_list, "GET", nil)
		json.Unmarshal(bytes, &cms)

		fmt.Println(&cms)
	}
}
