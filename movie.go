// Get all Movies
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/Movies", getMovies).Methods("GET")           //list of movie
	router.HandleFunc("/Movies/{id}", getMovie).Methods("GET")       //detail of movie
	router.HandleFunc("/Movies", createMovie).Methods("POST")        //add new movie
	router.HandleFunc("/Movies/{id}", updateMovie).Methods("PUT")    //update movie
	router.HandleFunc("/Movies/{id}", deleteMovie).Methods("DELETE") //delete movie

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

// Get a single Movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	MovieID, _ := strconv.Atoi(params["id"])
	for _, Movie := range Movies {
		if Movie.ID == MovieID {
			json.NewEncoder(w).Encode(Movie)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Create a new Movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Movie Movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.ID = len(Movies) + 1
	Movies = append(Movies, Movie)
	json.NewEncoder(w).Encode(Movie)
}

// Update a Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	MovieID, _ := strconv.Atoi(params["id"])
	for index, Movie := range Movies {
		if Movie.ID == MovieID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = MovieID
			Movies = append(Movies, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Delete a Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	MovieID, _ := strconv.Atoi(params["id"])
	for index, Movie := range Movies {
		if Movie.ID == MovieID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(nil)
}
