package db

import (
	"github.com/kylemadkins/leaderboard/models"
)

func (db *Database) GetLeaderboard() (*models.Leaderboard, error) {
	leaderboard := &models.Leaderboard{}
	leaderboard.Entries = make([]models.Entry, 0)
	query := `
		SELECT s.player_id, p.username, high_score, s.created_at as achieved_at
		FROM scores s
		JOIN players p ON s.player_id = p.id
		JOIN (
			SELECT player_id, MAX(score) AS high_score
			FROM scores
			GROUP BY player_id
		) high_scores ON s.player_id = high_scores.player_id AND s.score = high_score
		ORDER BY high_score DESC, achieved_at ASC
	`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return leaderboard, err
	}
	for rows.Next() {
		var entry models.Entry
		err := rows.Scan(&entry.PlayerID, &entry.Username, &entry.Score, &entry.AchievedAt)
		if err != nil {
			return leaderboard, err
		}
		leaderboard.Entries = append(leaderboard.Entries, entry)
	}
	return leaderboard, nil
}
