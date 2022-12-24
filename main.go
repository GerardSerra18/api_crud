package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	//Here we are initializing the new router
	r := mux.NewRouter()

	//Register the CRUD Handlers for Movies
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", readMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//Register the CRUD Handlers for Actors
	r.HandleFunc("/actors", createActor).Methods("POST")
	r.HandleFunc("/actors/{id}", readActor).Methods("GET")
	r.HandleFunc("/actors/{id}", updateActor).Methods("PUT")
	r.HandleFunc("/actors/{id}", deleteActor).Methods("DELETE")

	//Starting our server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
