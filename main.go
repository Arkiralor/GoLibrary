package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Initialize Mock Books
var books []Book

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Initialize Mock Authors
var authors []Author

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp := "Hello World!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetSingleBook(w http.ResponseWriter, r *http.Request) {

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Initialize the mux router
	router := mux.NewRouter()

	//Mock Data
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "1123581321",
			Title: "Wuthering Heights",
			Author: &Author{
				ID:        "1",
				FirstName: "Hugo",
				LastName:  "Chavez",
			},
		},
		Book{
			ID:    "2",
			Isbn:  "1129881321",
			Title: "Leroy Jenkins",
			Author: &Author{
				ID:        "2",
				FirstName: "Damien",
				LastName:  "Martinez",
			},
		},
		Book{
			ID:    "3",
			Isbn:  "1763581321",
			Title: "Henderson Economics of the Sorrowful State",
			Author: &Author{
				ID:        "3",
				FirstName: "Hector",
				LastName:  "Salmanenca",
			},
		},
	)

	// Route Handlers/Endpoints
	router.HandleFunc("/", HelloWorld).Methods("GET", "POST", "PUT", "PATCH", "DELETE")
	router.HandleFunc("/api/book/get/all", GetAllBooks).Methods("GET")
	router.HandleFunc("/api/book/get/{id:int}", GetSingleBook).Methods("GET")
	router.HandleFunc("/api/book/create/all", CreateBook).Methods("POST")
	router.HandleFunc("/api/book/update/all", UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/delete/all", DeleteBook).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
