package controller

import (
	book "BookStore/models"
	"BookStore/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook book.Book

var GetAllBook = func(w http.ResponseWriter, r *http.Request) {
	newBooks := book.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Write(res)
}

var GetBookById = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var id = params["bookId"]

	Id, err := strconv.ParseInt(id, 0, 0) // string이기 떄문에 int로 변환

	if err != nil {
		panic(err)
	}

	book, _ := book.GetAllBookById(Id)
	res, _ := json.Marshal(book)

	w.Write(res)
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	newBook := &book.Book{} // type를 Book으로 선언
	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Write(res)
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	updateBookId := mux.Vars(r)["bookId"]

	updateBookDate := &book.Book{}
	utils.ParseBody(r, updateBookDate)

	id, err := strconv.ParseInt(updateBookId, 0, 0)

	if err != nil {
		panic(err)
	}

	book, db := book.GetAllBookById(id)

	if updateBookDate.Name != "" {
		book.Name = updateBookDate.Name
	}

	if updateBookDate.Author != "" {
		book.Author = updateBookDate.Author
	}

	if updateBookDate.Publication != "" {
		book.Publication = updateBookDate.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Write(res)
}

var DeleteBookById = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var id = params["bookId"]

	Id, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		panic(err)
	}

	deletedBook, _ := book.GetAllBookById(Id)

	book.DeleteBook(Id)

	res, _ := json.Marshal(deletedBook)

	w.Write(res)
}
