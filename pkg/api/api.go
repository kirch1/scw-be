package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()

	//add middleware, in this case store DB to use later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/sets", func(r chi.Router) {
		r.Get("/", getSets)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func getSets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sets"))
}
