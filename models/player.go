package models

import (
	"fmt"
	"net/http"
	"regexp"
)

type Player struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type PlayerList struct {
	Players []Player `json:"players"`
}

func (p *Player) Bind(r *http.Request) error {
	if p.Username == "" {
		return fmt.Errorf("Username is a required field")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(p.Username) {
		return fmt.Errorf("Username must be alphanumeric")
	}
	return nil
}

func (*PlayerList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Player) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
