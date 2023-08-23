package api

import (
	"encoding/json"
	"log"
	"net/http"
	"scw-be/pkg/db/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-pg/pg/v10"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()

	//cors options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // List of allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum cache age in seconds
	})
	r.Use(corsOptions.Handler)

	//add middleware, in this case store DB to use later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/sets/{year}", func(r chi.Router) {
		r.Get("/", getSets)
	})

	r.Route("/sets", func(r chi.Router) {
		r.Get("/", getSets)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

type SetsResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Sets    []*models.Set `json:"sets"`
}

func getSets(w http.ResponseWriter, r *http.Request) {

	year := chi.URLParam(r, "year")
	log.Print(year)

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &SetsResponse{
			Success: false,
			Error:   "Could not get DB from context",
			Sets:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sets, err := models.GetSets(pgdb, year)
	if err != nil {
		res := &SetsResponse{
			Success: false,
			Error:   err.Error(),
			Sets:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &SetsResponse{
		Success: true,
		Error:   "",
		Sets:    sets,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding sets: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
