package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json: ID`
	ISBN     string    `json: ISBN`
	Name     string    `json: Name`
	Director *Director `json: Director`
}

type Director struct {
	FirstName string `json: FirstName`
	LastName  string `json: LastName`
}

var movies []Movies

func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := mux.Vars(r)
	for index, values := range movies {
		if values.ID == request["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := mux.Vars(r)

	for _, value := range movies {
		if value.ID == request["ID"] {
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	request := mux.Vars(r)
	for index, value := range movies {
		if value.ID == request["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = request["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	movies = append(movies, Movies{ID: "101", ISBN: "111111", Name: "John Wick", Director: &Director{FirstName: "Chad", LastName: "Stehalski"}})
	movies = append(movies, Movies{ID: "102", ISBN: "222222", Name: "Keanu Reeves", Director: &Director{FirstName: "Andrew Ross", LastName: "Sorkin"}})

	mux := mux.NewRouter()
	mux.HandleFunc("/", GetAllMovie).Methods("GET")
	mux.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	mux.HandleFunc("/movies", CreateMovie).Methods("POST")
	mux.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	mux.HandleFunc("/movies", DeleteMovie).Methods("DELETE")

	fmt.Printf("Listening at Port 8083...")
	http.ListenAndServe(":8083", mux)

}
