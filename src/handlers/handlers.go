package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marcosgfrites/REST-API_Golang/src/middlewares"
	"github.com/marcosgfrites/REST-API_Golang/src/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Handlers sets port, handler and do listening at port specified
func Handlers() {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/sign-up", middlewares.CheckDatabase(routers.Register)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//TODO: this permission will be changed. Now it allows all access
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
