package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kylemadkins/leaderboard/models"
)

func players(router chi.Router) {
	router.Get("/", getAllPlayers)
	router.Post("/", createPlayer)
}

func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := database.GetAllPlayers()
	if err != nil {
		render.Render(w, r, ServerError(err))
		return
	}
	if err := render.Render(w, r, players); err != nil {
		render.Render(w, r, ServerError(err))
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	player := &models.Player{}
	if err := render.Bind(r, player); err != nil {
		render.Render(w, r, BadRequest(err))
		return
	}
	if err := database.CreatePlayer(player); err != nil {
		render.Render(w, r, ServerError(err))
		return
	}
	if err := render.Render(w, r, player); err != nil {
		render.Render(w, r, ServerError(err))
	}
}
