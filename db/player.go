package db

import (
	"github.com/kylemadkins/leaderboard/models"
)

func (db *Database) GetAllPlayers() (*models.PlayerList, error) {
	list := &models.PlayerList{}
	list.Players = make([]models.Player, 0)
	rows, err := db.Conn.Query("SELECT * FROM players ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.ID, &player.Username, &player.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Players = append(list.Players, player)
	}
	return list, nil
}
