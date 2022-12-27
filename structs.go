package main

type Movie struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Actors []Actor
	Rating float64 `json:"rating"`
}

type Actor struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Movies      []Movie
	AudienceRtg float64 `json:"audience_rating"`
}

type MoviesActorsRelation struct {
	ActorID     int     `json:"actor_id"`
	MovieID     int     `json:"movie_id"`
	AudienceRtg float64 `json:"audience_rating"`
}
