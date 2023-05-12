package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/you/go-crud-api-test/pkg/handlers"
	"example.com/you/go-crud-api-test/pkg/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	movies := []models.Movie{
		{ID: "1", Isbn: "438227", Title: "Movie One", Director: &models.Director{Firstname: "John", Lastname: "Doe"}},
		{ID: "2", Isbn: "438312", Title: "Movie Two", Director: &models.Director{Firstname: "Steven", Lastname: "Smith"}},
	}
	handler := handlers.MovieHandler{Movies: movies}

	r.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
