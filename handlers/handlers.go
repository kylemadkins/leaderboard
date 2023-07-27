package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kylemadkins/leaderboard/db"
)

var database db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	database = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/players", players)
	router.Route("/scores", scores)
	router.Route("/leaderboard", leaderboard)
	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, ErrNotFound)
}
