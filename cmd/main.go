package main

import (
	"kvasir/pkg/handlers"
	"kvasir/pkg/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := repository.NewInMemoryStorage()
	manager := handlers.NewManager(db)
	r := chi.NewRouter()

	//r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	// POST /phrase
	r.Post("/phrase", manager.AddWord)
	// GET /phrase/list?limit=20&offset=0
	r.Get("/phrase/list", manager.GetDictionary)
	// POST /review
	r.Post("/review", manager.MakeReview)
	http.ListenAndServe(":3000", r)
}
