package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/siktro/books-api/platform/book"
	"github.com/siktro/books-api/platform/web"
)

type Book struct {
	DB     *sql.DB
	Logger *log.Logger
}

func (b *Book) GetBooks(w http.ResponseWriter, r *http.Request) error {
	list, err := book.List(r.Context(), b.DB)
	if err != nil {
		return err
	}

	return web.Respond(w, list, http.StatusOK)
}

// func (b *Book) GetBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	book := models.Book{}

// 	row := b.DB.QueryRow("SELECT * FROM books WHERE id = $1", params["id"])
// 	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			w.Write([]byte("no rows"))
// 		}
// 		b.Logger.Printf("scanning row", err)
// 	}

// 	if err = json.NewEncoder(w).Encode(book); err != nil {
// 		b.Logger.Printf("encoding json", err)
// 	}
// }

// func (b *Book) AddBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book

// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		b.Logger.Printf("decoding json", err)
// 	}

// 	row := b.DB.QueryRow(`
// 		INSERT INTO books(id, title, author, year)
// 		VALUES(default, $1, $2, $3)
// 		RETURNING id`,
// 		book.Title, book.Author, book.Year,
// 	)
// 	err = row.Scan(&book.ID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			w.Write([]byte("no rows"))
// 		} else {
// 			b.Logger.Printf("querying db", err)
// 		}
// 	}

// 	json.NewEncoder(w).Encode(book.ID)
// }

// func (b *Book) UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book

// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		b.Logger.Printf("decoding json", err)
// 	}

// 	res, err := b.DB.Exec(`
// 		UPDATE books SET
// 			title = $1,
// 			author = $2,
// 			year = $3
// 		WHERE
// 			id = $4
// 		RETURNING id
// 	`, &book.Title, &book.Author, &book.Year, &book.ID)

// 	if err != nil {
// 		b.Logger.Printf("querying db", err)
// 	}

// 	rf, err := res.RowsAffected()
// 	if err != nil {
// 		b.Logger.Printf("querying db", err)
// 	}

// 	json.NewEncoder(w).Encode(rf)
// }

// func (b *Book) RemoveBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	res, err := b.DB.Exec(`
// 		DELETE FROM books
// 		WHERE id = $1
// 	`, params["id"])

// 	if err != nil {
// 		b.Logger.Printf("querying db", err)
// 	}

// 	rf, err := res.RowsAffected()
// 	if err != nil {
// 		b.Logger.Printf("querying db", err)
// 	}

// 	json.NewEncoder(w).Encode(rf)
// }
