package main

import "database/sql"

func initDB() error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE movies (id SERIAL PRIMARY KEY, title TEXT, year INTEGER, genre TEXT, rating REAL)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE actors (id SERIAL PRIMARY KEY, first_name TEXT, last_name TEXT, gender CHAR(1), age INTEGER)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE movies_actors_relation (actor_id INTEGER REFERENCES actors(id), movie_id INTEGER REFERENCES movies(id), audience_rating REAL, PRIMARY KEY (actor_id, movie_id))`)
	if err != nil {
		return err
	}

	return nil
}

// Database connection
func connectToDB() (*sql.DB, error) {
	//psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, password, dbname)
	connStr := "user=gerard password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
