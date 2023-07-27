package models

import (
	"net/http"
)

type Score struct {
	ID        int    `json:"id"`
	PlayerID  int    `json:"player_id"`
	Score     int    `json:"score"`
	CreatedAt string `json:"created_at"`
}

type ScoreList struct {
	Scores []Score `json:"scores"`
}

func (p *Score) Bind(r *http.Request) error {
	return nil
}

func (*ScoreList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Score) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
