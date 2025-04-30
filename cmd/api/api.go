package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/handler"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/repository"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/service"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/pkg/db"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(2 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	db := db.GetDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("List of users"))
			})
			r.Post("/", userHandler.CreateUser)
		})

	})

	return r

}

func (app *application) run(mux http.Handler) error {

	server := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		IdleTimeout:  time.Minute * 1,
	}

	fmt.Println("Starting server on port")

	return server.ListenAndServe()

}
