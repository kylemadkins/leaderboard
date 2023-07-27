package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func leaderboard(router chi.Router) {
	router.Get("/", getLeaderboard)
}

func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboard, err := database.GetLeaderboard()
	if err != nil {
		render.Render(w, r, ServerError(err))
	}
	if err := render.Render(w, r, leaderboard); err != nil {
		render.Render(w, r, ServerError(err))
	}
}
