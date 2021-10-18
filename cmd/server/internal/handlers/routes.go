package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siktro/books-api/platform/web"
)

func MakeAPI(db *sql.DB, logger *log.Logger) http.Handler {

	app := web.NewApp(logger, mux.NewRouter())

	// Book handlers.
	b := Book{DB: db, Logger: logger}

	app.Handle("/books", b.GetBooks, "GET")
	// router.HandleFunc("/books", app.Handle(b.GetBooks)).Methods("GET")
	// router.HandleFunc("/books/{id}", b.GetBook).Methods("GET")
	// router.HandleFunc("/books", b.AddBook).Methods("POST")
	// router.HandleFunc("/books", b.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", b.RemoveBook).Methods("DELETE")

	// TODO: middleware expample
	// router.Use(func(next http.Handler) http.Handler {
	// 	fn := func(w http.ResponseWriter, r *http.Request) {
	// 		next.ServeHTTP(w, r)
	// 		log.Println("mid1")
	// 	}
	// 	return http.HandlerFunc(fn)
	// })

	return app
}
