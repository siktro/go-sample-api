package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/siktro/books-api/database"
	"github.com/siktro/books-api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	// Setup env. variables.
	err := setEnvFromFile("./.env")
	logFatal("reading .env file", err)

	// DB connetction.
	var closer func()
	db, closer, err = database.Open(&database.Config{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	})
	logFatal("opening database", err)
	defer closer()

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updBook).Methods("PUT")
	router.HandleFunc("/books/{id}", delBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	book := models.Book{}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		logFatal("reporting to client", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logFatal("reporting to client", err)
		}
		books = append(books, book)
	}

	if err = json.NewEncoder(w).Encode(books); err != nil {
		logFatal("encoding json", err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := models.Book{}

	row := db.QueryRow("SELECT * FROM books WHERE id = $1", params["id"])
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.Write([]byte("no rows"))
		}
		logFatal("scanning row", err)
	}

	if err = json.NewEncoder(w).Encode(book); err != nil {
		logFatal("encoding json", err)
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		logFatal("decoding json", err)
	}

	row := db.QueryRow(`
		INSERT INTO books(id, title, author, year)
		VALUES(default, $1, $2, $3)
		RETURNING id`,
		book.Title, book.Author, book.Year,
	)
	err = row.Scan(&book.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Write([]byte("no rows"))
		} else {
			logFatal("querying db", err)
		}
	}

	json.NewEncoder(w).Encode(book.ID)
}

func updBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		logFatal("decoding json", err)
	}

	res, err := db.Exec(`
		UPDATE books SET
			title = $1,
			author = $2,
			year = $3
		WHERE
			id = $4
		RETURNING id
	`, &book.Title, &book.Author, &book.Year, &book.ID)

	if err != nil {
		logFatal("querying db", err)
	}

	rf, err := res.RowsAffected()
	if err != nil {
		logFatal("querying db", err)
	}

	json.NewEncoder(w).Encode(rf)
}

func delBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := db.Exec(`
		DELETE FROM books
		WHERE id = $1
	`, params["id"])

	if err != nil {
		logFatal("querying db", err)
	}

	rf, err := res.RowsAffected()
	if err != nil {
		logFatal("querying db", err)
	}

	json.NewEncoder(w).Encode(rf)
}

// Utils.

func setEnvFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading file: %w;", err)
	}

	reg, err := regexp.Compile(`(\w+)\s*=\s*(.+)`)
	if err != nil {
		return fmt.Errorf("regexp compilation: %w;", err)
	}

	pairs := reg.FindAllSubmatch(data, -1)
	for _, pair := range pairs {
		key := string(pair[1])
		value := string(pair[2])
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("setting envvar [%s=%s]: %w", key, value, err)
		}
	}

	return nil
}

func logFatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
