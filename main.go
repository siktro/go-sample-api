package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/siktro/books-api/controllers"
	"github.com/siktro/books-api/database"
)

var db *sql.DB

func main() {
	logger := log.New(os.Stdout, "[>] ", log.LstdFlags)

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

	api := controllers.MakeAPI(db, logger)
	log.Fatal(http.ListenAndServe(":8000", api))
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
