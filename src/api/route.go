package api

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		{
			Id:    1,
			Title: "Title 1",
			Text:  "Text 1",
		},
		{
			Id:    2,
			Title: "Title 2",
			Text:  "Text 2",
		},
		{
			Id:    3,
			Title: "Title 3",
			Text:  "Text 3",
		},
	}
}

func GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	response.WriteHeader(http.StatusOK)

	response.Write(result)
}

func AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newPost Post
	err := json.NewDecoder(request.Body).Decode(&newPost)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	newPost.Id = len(posts) + 1
	posts = append(posts, newPost)

	response.WriteHeader(http.StatusOK)

	result, err := json.Marshal(newPost)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	response.Write(result)
}
