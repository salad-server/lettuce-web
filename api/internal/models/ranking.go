package models

import (
	"context"
	"fmt"
	"log"
	"mini-api/internal/scores"
	"time"
)

type PlayerRanking struct {
	PP          float32 `json:"pp"`
	Accuracy    float32 `json:"acc"`
	TotalScore  int     `json:"total_score"`
	RankedScore int     `json:"ranked_score"`
	Plays       int     `json:"playcount"`
	User        struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Country  string `json:"country"`
	} `json:"user"`
}

func (db *DB) Leaderboard(pp bool, mode string, page int) ([]PlayerRanking, error) {
	m := scores.Modes[mode]
	o := "s.rscore"

	if pp {
		o = "s.pp"
	}

	q := fmt.Sprintf(`
		SELECT
			u.id, u.name, u.country, s.pp,
			s.rscore, s.tscore, s.plays, s.acc
		FROM users u
		JOIN stats s ON u.id = s.id
		WHERE s.mode = ?
		ORDER BY %s DESC
		LIMIT ?, 10
	`, o)

	var ranking []PlayerRanking
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, m, page*PAGE_LEN)

	defer cancel()

	if err != nil {
		log.Println("Error in Leaderboard")
		return nil, err
	}

	for row.Next() {
		var rank PlayerRanking

		if err := row.Scan(
			&rank.User.ID,
			&rank.User.Username,
			&rank.User.Country,
			&rank.PP,
			&rank.RankedScore,
			&rank.TotalScore,
			&rank.Plays,
			&rank.Accuracy,
		); err != nil {
			log.Println("Error in Leaderboard")
			return nil, err
		}

		ranking = append(ranking, rank)
	}

	return ranking, nil
}
