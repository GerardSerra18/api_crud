package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	// router initalization
	r := mux.NewRouter()

	// endpoints handlers for movie API

	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", readMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", getAllMovies).Methods("GET")

	// endpoints handlers for actor API
	r.HandleFunc("/actors", createActor).Methods("POST")
	r.HandleFunc("/actors/{id}", readActor).Methods("GET")
	r.HandleFunc("/actors/{id}", updateActor).Methods("PUT")
	r.HandleFunc("/actors/{id}", deleteActor).Methods("DELETE")
	r.HandleFunc("/actors", getAllActors).Methods("GET")
	r.HandleFunc("/actors/{id}/rating", getActorRating).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
