//done with postman

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

type Movie struct { //creating a movies struct to store all movies
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct { //movie directors struct
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) { //pointer of postman request passed here as "r".
	w.Header().Set("Content-Type", "application/json") //w is the writer, he sends response back to postman
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //were deleting movie and adding new movie with updated value
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "10000", Title: "Movie1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "10001", Title: "Movie2", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})
	movies = append(movies, Movie{"3", "10002", "Movie3", &Director{"Trisha", "Songra"}})
	//we wish to populate now with some values so it wont look empty in the start

	//important to note that the paths given below are all ROUTES
	r.HandleFunc("/movies", getMovies).Methods("GET")           //handlefunc for get all movies
	r.HandleFunc("/movies/id", getMovie).Methods("GET")         //handlefunc for get single movie by ID
	r.HandleFunc("/movies", createMovie).Methods("POST")        //handle func for creating movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    //handle func for updating movie by ID
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") //handle func for deleting movie by id

	fmt.Printf("starting server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
