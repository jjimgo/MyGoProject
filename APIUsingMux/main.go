package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"Id"`
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"Title"`
	Author *Author `json:"Auther"`
}

type Author struct {
	FirstName string `json:"FirstName`
	LastName  string `json:"LastName`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal(books) //

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("id")

	params := mux.Vars(r)
	fmt.Println(params)

	for _, item := range books {
		if item.ID == params["id"] {
			res, _ := json.Marshal(item)
			w.Write(res)
			return
		}
	}

	fmt.Fprint(w, "none Data")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	r.ParseForm()

	fmt.Println(r.Body)

	_ = json.NewDecoder(r.Body).Decode(&book)

	fmt.Println(book)

	book.ID = "1231234"

	books = append(books, book)

	res, _ := json.Marshal(book)

	w.Write(res)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)

			res, _ := json.Marshal(books)
			w.Write(res)
			break
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	res, _ := json.Marshal(books)

	w.Write(res)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "1", Title: "myFirstBook", Author: &Author{FirstName: "Yu", LastName: "Jin"}})

	books = append(books, Book{ID: "2", Isbn: "2", Title: "mySecondBook", Author: &Author{FirstName: "Ho", LastName: "Jin"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("PUT")
	r.HandleFunc("/api/books", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
