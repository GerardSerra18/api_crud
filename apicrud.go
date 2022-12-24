package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

//CRUD operations for movies

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the movie data
	var newMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the data
	if newMovie.Title == "" {
		http.Error(w, "Missing title", http.StatusBadRequest)
		return
	}
	if newMovie.Year.IsZero() {
		http.Error(w, "Missing Year", http.StatusBadRequest)
		return
	}

	// Save the movie to the database
	id, err := saveMovieToDB(newMovie)
	if err != nil {
		http.Error(w, "Error saving movie to database", http.StatusInternalServerError)
		return
	}

	// Return a response with the newly created movie
	newMovie.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func readMovie(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the movie ID
	params := mux.Vars(r)
	id := params["id"]

	// Retrieve the movie from the database
	movie, err := readMovieFromDB(id)
	if err != nil {
		http.Error(w, "Error retrieving movie from database", http.StatusInternalServerError)
		return
	}

	// Return a response with the movie data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the movie ID
	params := mux.Vars(r)
	id := params["id"]

	// Parse the request body to get the updated movie data
	var updatedMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the data
	if updatedMovie.Title == "" {
		http.Error(w, "Missing title", http.StatusBadRequest)
		return
	}
	if updatedMovie.Year.IsZero() {
		http.Error(w, "Missing release date", http.StatusBadRequest)
		return
	}

	// Update the movie in the database
	if err := updateMovieInDB(id, updatedMovie); err != nil {
		// Return a response with the updated movie
		updatedMovie.ID = id
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedMovie)
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the movie ID
	params := mux.Vars(r)
	id := params["id"]

	// Delete the movie from the database
	if err := deleteMovieFromDB(id); err != nil {
		http.Error(w, "Error deleting movie from database", http.StatusInternalServerError)
		return
	}

	// Return a response with an HTTP status code indicating success
	w.WriteHeader(http.StatusNoContent)
}

//CRUD operations for actors

func createActor(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the actor data
	var newActor Actor
	if err := json.NewDecoder(r.Body).Decode(&newActor); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the data
	if newActor.FirstName == "" {
		http.Error(w, "Missing first name", http.StatusBadRequest)
		return
	}
	if newActor.LastName == "" {
		http.Error(w, "Missing last name", http.StatusBadRequest)
		return
	}
	if newActor.Gender == "" {
		http.Error(w, "Missing gender", http.StatusBadRequest)
		return
	}
	if newActor.Age == 0 {
		http.Error(w, "Missing age", http.StatusBadRequest)
		return
	}

	// Save the actor to the database
	id, err := saveActorToDB(newActor)
	if err != nil {
		http.Error(w, "Error saving Actor to database", http.StatusInternalServerError)
		return
	}
	// Return a response with the newly created actor
	newActor.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newActor)
}

func readActor(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the actor ID
	params := mux.Vars(r)
	id := params["id"]

	// Retrieve the actor from the database
	actor, err := readActorFromDB(id)
	if err != nil {
		http.Error(w, "Error retrieving actor from database", http.StatusInternalServerError)
		return
	}

	// Return a response with the actor data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actor)
}

func updateActor(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the actor ID
	params := mux.Vars(r)
	id := params["id"]

	// Parse the request body to get the updated actor data
	var updatedActor Actor
	if err := json.NewDecoder(r.Body).Decode(&updatedActor); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the data
	if updatedActor.FirstName == "" {
		http.Error(w, "Missing first name", http.StatusBadRequest)
		return
	}
	if updatedActor.LastName == "" {
		http.Error(w, "Missing last name", http.StatusBadRequest)
		return
	}
	if updatedActor.Gender == "" {
		http.Error(w, "Missing gender", http.StatusBadRequest)
		return
	}
	if updatedActor.Age == 0 {
		http.Error(w, "Missing age", http.StatusBadRequest)
		return
	}

	// Update the actor in the database
	if err := updateActorInDB(id, updatedActor); err != nil {
		http.Error(w, "Error updating actor in database", http.StatusInternalServerError)
		return
	}

	// Return a response with the updated actor
	updatedActor.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedActor)
}

func deleteActor(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters to get the actor ID
	params := mux.Vars(r)
	id := params["id"]

	// Delete the actor from the database
	if err := deleteActorFromDB(id); err != nil {
		http.Error(w, "Error deleting actor from database", http.StatusInternalServerError)
		return
	}

	// Return a response with an HTTP status code indicating success
	w.WriteHeader(http.StatusNoContent)
}


