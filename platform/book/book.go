package book

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

func List(ctx context.Context, db *sql.DB) ([]Book, error) {
	var books []Book
	book := Book{}

	rows, err := db.QueryContext(ctx, `
		SELECT
			*
		FROM
			books
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func Get(ctx context.Context, db *sql.DB, id string) (*Book, error) {
	book := Book{}

	row := db.QueryRow(`
		SELECT
			*
		FROM
			books
		WHERE id = $1
	`, id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &book, nil
}

func Put(ctx context.Context, db *sql.DB) (int, error) {
	var book Book

	row := db.QueryRow(`
		INSERT INTO
			books(id, title, author, year)
		VALUES
			(default, $1, $2, $3)
		RETURNING id
		`, book.Title, book.Author, book.Year,
	)

	err := row.Scan(&book.ID)
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}
