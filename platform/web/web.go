package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

type App struct {
	logger *log.Logger
	router *mux.Router
}

func NewApp(logger *log.Logger, router *mux.Router) *App {
	return &App{
		logger: logger,
		router: router,
	}
}

func (a *App) Handle(pattern string, handler HandlerWithError, methods ...string) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			a.logger.Printf("ERROR : %v\n", err)

			err = RespondError(w, err)
			if err != nil {
				a.logger.Printf("ERROR : %v\n", err)
			}
		}
	}

	a.router.HandleFunc(pattern, fn).Methods(methods...)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
