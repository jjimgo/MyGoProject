package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	router "BookStore/routes"
)

func main() {
	r := mux.NewRouter()

	router.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
