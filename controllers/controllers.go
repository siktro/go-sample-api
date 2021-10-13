package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siktro/books-api/controllers/handlers"
)

// TODO: for future
// type extendedHandler func(http.ResponseWriter, *http.Request) error

// func handle(fn extendedHandler) http.HandlerFunc {
// 	closure := func(w http.ResponseWriter, r *http.Request) {
// 		if err := fn(w, r); err != nil {

// 		}
// 	}

// 	return closure
// }

func MakeAPI(db *sql.DB, logger *log.Logger) http.Handler {

	var router = mux.NewRouter()

	// Book handlers.
	b := handlers.Book{DB: db, Logger: logger}

	router.HandleFunc("/books", b.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", b.GetBook).Methods("GET")
	router.HandleFunc("/books", b.AddBook).Methods("POST")
	router.HandleFunc("/books", b.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", b.RemoveBook).Methods("DELETE")

	// TODO: is it run before or after?
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.RequestURI)
			})
	})

	return router
}
