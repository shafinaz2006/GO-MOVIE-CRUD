package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// json.NewEncoder(w).Encode(movies) // this will return the other movies.
			json.NewEncoder(w).Encode("Movie deleted") // this will return a string
			break
		}
	}
	json.NewEncoder(w).Encode("Movie ID not found")
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode("Movie ID not found")
}

func createMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = params["id"]
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
	json.NewEncoder(w).Encode("Movie ID not found")
}

func main() {
	fmt.Println("GO CRUD MOVIE APP")
	router := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "2", Title: "one", Director: &Director{Firstname: "John", Lastname: "DD"}})
	movies = append(movies, Movie{ID: "133", Isbn: "2232", Title: "Two", Director: &Director{Firstname: "Jane", Lastname: "Smith"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
