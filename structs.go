package main

import "time"

type Movie struct {
	ID     int64      `json:"id"`
	Title  string     `json:"title"`
	Year   *time.Time `json:"year"`
	Actors []Actor    `json:"actors"`
	Genre  string     `json:"gen"`
	Rating int64      `json:"rating"`
}

type Actor struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    string  `json:"gender"`
	Age       int64   `json:"age"`
	Movies    []Movie `json:"movies"`
}
