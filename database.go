package main

import (
	"database/sql"
	"strconv"
)

func saveMovieToDB(movie Movie) (string, error) {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Insert the movie data into the database
	res, err := db.Exec("INSERT INTO movies (title, release_date, actors, genre, rating) VALUES (?, ?, ?, ?, ?)", movie.Title, movie.Year, movie.Actors, movie.Genre, movie.Rating)
	if err != nil {
		return "", err
	}

	// Get the ID of the newly created record
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func readMovieFromDB(id string) (Movie, error) {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return Movie{}, err
	}
	defer db.Close()

	// Retrieve the movie from the database
	var m Movie
	err = db.QueryRow("SELECT id, title, release_date, actors, genre, rating FROM movies WHERE id = ?", id).Scan(&m.ID, &m.Title, &m.Year, &m.Actors, &m.Genre, &m.Rating)
	if err != nil {
		return Movie{}, err
	}

	return m, nil
}
func updateMovieInDB(id string, movie Movie) error {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Update the movie in the database
	_, err = db.Exec("UPDATE movies SET title = ?, release_date = ?, actors = ?, genre = ?, rating = ? WHERE id = ?", movie.Title, movie.Year, movie.Actors, movie.Genre, movie.Rating, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteMovieFromDB(id string) error {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete the movie from the database
	_, err = db.Exec("DELETE FROM movies WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func saveActorToDB(actor Actor) (string, error) {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Insert the actor data into the database
	res, err := db.Exec("INSERT INTO actors (first_name, last_name, gender, age, movies) VALUES (?, ?, ?, ?, ?)", actor.FirstName, actor.LastName, actor.Gender, actor.Age, actor.Movies)
	if err != nil {
		return "", err
	}

	// Get the ID of the newly created record
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func readActorFromDB(id string) (Actor, error) {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return Actor{}, err
	}
	defer db.Close()

	// Retrieve the actor from the database
	var a Actor
	err = db.QueryRow("SELECT id, first_name, last_name, gender, age, movies FROM actors WHERE id = ?", id).Scan(&a.ID, &a.FirstName, &a.LastName, &a.Gender, &a.Age, &a.Movies)
	if err != nil {
		return Actor{}, err
	}

	return a, nil
}

func updateActorInDB(id string, actor Actor) error {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Update the actor in the database
	_, err = db.Exec("UPDATE actors SET first_name = ?, last_name = ?, gender = ?, age = ?, movies = ? WHERE id = ?", actor.FirstName, actor.LastName, actor.Gender, actor.Age, actor.Movies, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteActorFromDB(id string) error {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete the actor from the database
	_, err = db.Exec("DELETE FROM actors WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

