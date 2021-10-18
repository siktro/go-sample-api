# Simple Web API written in Go

For practice and learning sake.

> `.env` is a dummy file with no real credentials.

## Endpoints

- `GET /books` - get all books;
- `GET /books/{id}` - get a book by id;
- `POST /books` - add a book;
- `PUT /books` - update a book;
- `DELETE /books/{id}` - delete a book by id.

# Run

1. `docker compose up -d` - run DB with GUI.
2. `go run main.go`

# TODO

- [ ] Start containers from server itself?
- [ ] Add some middleware.
- [ ] Add auth.
- [ ] ORM?
- [ ] Add some 'seed' commands to populate DB.
