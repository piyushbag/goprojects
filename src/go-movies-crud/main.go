package main

import (
	"fmt"
	"net/http"

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

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", GetAllMovies()).Methods("GET")
	mux.HandleFunc("/movies/{id}", GetMovie()).Methods("GET")
	mux.HandleFunc("/movies", CreateMovie()).Methods("POST")
	mux.HandleFunc("/movies/{id}", UpdateMovie()).Methods("PUT")
	mux.HandleFunc("/movies", DeleteMovie()).Methods("DELETE")

	fmt.Printf("Listening at Port 8083...")
	http.ListenAndServe(":8083", mux)

}
