package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kylemadkins/leaderboard/models"
)

func scores(router chi.Router) {
	router.Get("/", getAllScores)
	router.Post("/", createScore)
}

func getAllScores(w http.ResponseWriter, r *http.Request) {
	scores, err := database.GetAllScores()
	if err != nil {
		render.Render(w, r, ServerError(err))
		return
	}
	if err := render.Render(w, r, scores); err != nil {
		render.Render(w, r, ServerError((err)))
	}
}

func createScore(w http.ResponseWriter, r *http.Request) {
	score := &models.Score{}
	if err := render.Bind(r, score); err != nil {
		render.Render(w, r, BadRequest(err))
		return
	}
	if err := database.CreateScore(score); err != nil {
		render.Render(w, r, ServerError(err))
		return
	}
	if err := render.Render(w, r, score); err != nil {
		render.Render(w, r, ServerError(err))
	}
}
