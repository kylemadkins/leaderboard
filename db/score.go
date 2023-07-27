package db

import (
	"github.com/kylemadkins/leaderboard/models"
)

func (db *Database) GetAllScores() (*models.ScoreList, error) {
	list := &models.ScoreList{}
	list.Scores = make([]models.Score, 0)
	rows, err := db.Conn.Query("SELECT * FROM scores ORDER BY created_at DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var score models.Score
		err := rows.Scan(&score.ID, &score.PlayerID, &score.Score, &score.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Scores = append(list.Scores, score)
	}
	return list, nil
}

func (db *Database) CreateScore(score *models.Score) error {
	query := `INSERT INTO scores (player_id, score) VALUES ($1, $2) RETURNING id, player_id, score, created_at`
	err := db.Conn.QueryRow(query, score.PlayerID, score.Score).Scan(&score.ID, &score.PlayerID, &score.Score, &score.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
