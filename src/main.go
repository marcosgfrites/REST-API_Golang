package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marcosgfrites/REST-API_Golang/src/api"
	"log"
	"net/http"
)

const (
	port string = ":8080"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/",
		func(response http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(response, "Up and running...")
		},
	)

	router.HandleFunc("/posts", api.GetPosts).Methods("GET")

	router.HandleFunc("/posts", api.AddPost).Methods("POST")

	log.Println("Server listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
