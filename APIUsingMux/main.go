package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
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
	r.ParseForm()

	var bookSecond Book

	fmt.Println("before Json.NewDecoder", bookSecond) // 데이터가 없으니깐 = nil
	fmt.Println("r.Body Data", r.Body)                // body에 입력이 되어서 들어오는 Data

	_ = json.NewDecoder(r.Body).Decode(&bookSecond)

	fmt.Println("after Json NewDecoder", bookSecond) // type에 있는 데이터만 Decoder 된다.

	bookSecond.ID = "1231234"

	books = append(books, bookSecond)

	res, _ := json.Marshal(bookSecond)

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

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ParseForm 실행 전 r.Body")
	fmt.Println(r.Body)

	r.ParseForm()

	fmt.Println("ParseForm 실행 후 r.Body")
	fmt.Println(r.Body)

	res, _ := ioutil.ReadAll(r.Body)

	fmt.Println("ParseForm 실행 후, ioutil.ReadAll r.Body")
	fmt.Println(res)

	var book Book

	json.Unmarshal([]byte(res), book)

	fmt.Println("jsonUnmarshal r.Body")
	fmt.Println(book)

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

	r.HandleFunc("/test", test).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
