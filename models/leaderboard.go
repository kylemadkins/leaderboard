package models

import (
	"net/http"
)

type Entry struct {
	PlayerID   int    `json:"player_id"`
	Username   string `json:"username"`
	Score      int    `json:"score"`
	AchievedAt string `json:"achieved_at"`
}

type Leaderboard struct {
	Entries []Entry `json:"entries"`
}

func (l *Leaderboard) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (e *Entry) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
