package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
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
		http.Error(w, "Missing release date", http.StatusBadRequest)
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
	movie, err := getMovieFromDB(id)
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
