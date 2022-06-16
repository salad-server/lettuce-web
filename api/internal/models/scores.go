package models

import (
	"context"
	"fmt"
	"log"
	"time"
)

const PAGE_LEN = 10

type Score struct {
	ID       int     `json:"id"`
	MD5      string  `json:"md5"`
	Score    int     `json:"score"`
	PP       float32 `json:"pp"`
	Acc      float32 `json:"acc"`
	MaxCombo int     `json:"max_combo"`
	Mods     int     `json:"mods"`
	Count300 int     `json:"c300"`
	Count100 int     `json:"c100"`
	Count50  int     `json:"c50"`
	Geki     int     `json:"geki"`
	Katu     int     `json:"katu"`
	Miss     int     `json:"miss"`
	Grade    string  `json:"rank"`
	Status   int     `json:"submission_status"`
	Date     []uint8 `json:"play_date"`
	Perfect  bool    `json:"perfect_score"`
}

func (db *DB) GetScores(best bool, uid, mode, page int) ([]Score, error) {
	sort := "id"
	ranked := ""

	if best {
		sort = "pp"
		ranked = "AND m.status IN (2, 3, 4) AND s.status = 2"
	}

	q := fmt.Sprintf(`
		SELECT
			s.id, s.map_md5, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect 
		FROM scores s
		INNER JOIN maps m ON s.map_md5 = m.md5
		WHERE s.userid = ? AND s.mode = ? %s
		ORDER BY %s DESC LIMIT ?, 10
	`, ranked, sort)

	var scores []Score
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, uid, mode, page*PAGE_LEN)

	defer cancel()

	if err != nil {
		log.Println("Error in GetScores")
		return nil, err
	}

	for row.Next() {
		var score Score

		if err := row.Scan(
			&score.ID,
			&score.MD5,
			&score.Score,
			&score.PP,
			&score.Acc,
			&score.MaxCombo,
			&score.Mods,
			&score.Count300,
			&score.Count100,
			&score.Count50,
			&score.Miss,
			&score.Geki,
			&score.Katu,
			&score.Grade,
			&score.Status,
			&score.Date,
			&score.Perfect,
		); err != nil {
			log.Println("Error in GetScores")
			return nil, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}
