package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Handlers sets port, handler and do listening at port specified
func Handlers() {
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//TODO: this permission will be changed. Now it allows all access
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}