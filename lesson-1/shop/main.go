package main

import (
	"gb-go-architecture/lesson-1/shop/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(s *server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/item", s.listItemHandler).Methods("GET")
	router.HandleFunc("/item", s.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", s.updateItemHandler).Methods("PUT")

	return router
}

func main() {
	s := &server{
		rep: repository.NewMapDB(),
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: NewRouter(s),
	}
	log.Fatal(srv.ListenAndServe())
}
