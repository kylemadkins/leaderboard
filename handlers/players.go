package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func players(router chi.Router) {
	router.Get("/", getAllPlayers)
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
