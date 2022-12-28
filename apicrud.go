package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

//CRUD operations for movies

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate movie data

	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(`INSERT INTO movies (title, year, genre, rating) VALUES ($1, $2, $3, $4) RETURNING id`, movie.Title, movie.Year, movie.Genre, movie.Rating).Scan(&movie.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func readMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var movie Movie
	err = db.QueryRow(`SELECT id, title, year, genre, rating FROM movies WHERE id=$1`, id).Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Genre, &movie.Rating)
	if err == sql.ErrNoRows {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var movie Movie
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate movie data

	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE movies SET title=$1, year=$2, genre=$3, rating=$4 WHERE id=$5`, movie.Title, movie.Year, movie.Genre, movie.Rating, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM movies WHERE id=$1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CRUD operations for actors7
func createActor(w http.ResponseWriter, r *http.Request) {
	// parse request body
	var a Actor
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate input
	if a.FirstName == "" || a.LastName == "" || a.Gender == "" || a.Age == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute INSERT query
	var actorID int
	err = db.QueryRow("INSERT INTO actors (first_name, last_name, gender, age) VALUES ($1, $2, $3, $4) RETURNING id", a.FirstName, a.LastName, a.Gender, a.Age).Scan(&actorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.ID = actorID

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// serialize actor to JSON
	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func readActor(w http.ResponseWriter, r *http.Request) {
	// get actor ID from URL path
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute SELECT query
	var a Actor
	err = db.QueryRow("SELECT id, first_name, last_name, gender, age FROM actors WHERE id = $1", id).Scan(&a.ID, &a.FirstName, &a.LastName, &a.Gender, &a.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Actor not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write JSON response
	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func updateActor(w http.ResponseWriter, r *http.Request) {
	// get actor ID from URL path
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// parse request body
	var a Actor
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate input
	if a.FirstName == "" || a.LastName == "" || a.Gender == "" || a.Age == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute UPDATE query
	res, err := db.Exec("UPDATE actors SET first_name = $1, last_name = $2, gender = $3, age = $4 WHERE id = $5", a.FirstName, a.LastName, a.Gender, a.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if actor was updated
	n, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if n == 0 {
		http.Error(w, "Actor not found", http.StatusNotFound)
		return
	}

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write JSON response
	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteActor(w http.ResponseWriter, r *http.Request) {
	// get actor ID from URL path
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute DELETE query
	_, err = db.Exec("DELETE FROM actors WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response
	w.WriteHeader(http.StatusNoContent)
}

//GetAllMovies and GetAllActors and with rating ang pagination, BONUS PART

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	// parse pagination query parameters
	pageSizeStr := r.URL.Query().Get("page_size")
	pageNumberStr := r.URL.Query().Get("page_number")
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10 // default page size
	}
	pageNumber, err := strconv.ParseInt(pageNumberStr, 10, 64)
	if err != nil {
		pageNumber = 1 // default page number
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute SELECT query
	offset := (pageNumber - 1) * pageSize
	rows, err := db.Query("SELECT id, title, year, genre, rating FROM movies ORDER BY year DESC LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// iterate over rows
	var movies []Movie
	for rows.Next() {
		var m Movie
		err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.Genre, &m.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, m)
	}

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write JSON response
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllActors(w http.ResponseWriter, r *http.Request) {
	// parse pagination query parameters
	pageSizeStr := r.URL.Query().Get("page_size")
	pageNumberStr := r.URL.Query().Get("page_number")
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10 // default page size
	}
	pageNumber, err := strconv.ParseInt(pageNumberStr, 10, 64)
	if err != nil {
		pageNumber = 1 // default page number
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute SELECT query
	offset := (pageNumber - 1) * pageSize
	rows, err := db.Query("SELECT id, first_name, last_name, gender, age FROM actors ORDER BY first_name ASC LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// iterate over rows
	var actors []Actor
	for rows.Next() {
		var a Actor
		err := rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Gender, &a.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		actors = append(actors, a)
	}

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write JSON response
	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getActorRating(w http.ResponseWriter, r *http.Request) {
	// parse actor ID from URL path
	vars := mux.Vars(r)
	actorIDStr := vars["id"]
	actorID, err := strconv.ParseInt(actorIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// connect to database
	db, err := connectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// execute SELECT query
	var rating float64
	err = db.QueryRow("SELECT AVG(audience_rating) FROM movies_actors_relation WHERE actor_id = $1", actorID).Scan(&rating)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Actor not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write JSON response
	err = json.NewEncoder(w).Encode(struct {
		Rating float64 `json:"rating"`
	}{
		Rating: rating,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

