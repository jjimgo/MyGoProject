package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "helllo")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseFrom() Error", err)
		return
	}
	// Body를 통해서 받는 것이 아니라 Params를 통해서 Key - vaule형식으로 받는 것
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	fmt.Fprintf(w, "Post Request SuccessFul")
}

func main() {

	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer)
	r.HandleFunc("/form", formHandler).Methods("POST")
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	fmt.Printf("starting server at port 8080\n")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
