package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeAPI(db *sql.DB, logger *log.Logger) http.Handler {

	var router = mux.NewRouter()

	// Book handlers.
	b := Book{DB: db, Logger: logger}

	router.HandleFunc("/books", b.GetBooks).Methods("GET")
	// router.HandleFunc("/books/{id}", b.GetBook).Methods("GET")
	// router.HandleFunc("/books", b.AddBook).Methods("POST")
	// router.HandleFunc("/books", b.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", b.RemoveBook).Methods("DELETE")

	// TODO: is it run before or after?
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.RequestURI)
				next.ServeHTTP(w, r)
			})
	})

	router.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			log.Println("mid1")
		}
		return http.HandlerFunc(fn)
	})

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
				log.Println("mid2")
			})
	})

	return router
}
