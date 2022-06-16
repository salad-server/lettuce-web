package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const PAGE_LEN = 10

type Score struct {
	ID       int     `json:"id"`
	Score    int     `json:"score"`
	PP       float32 `json:"pp"`
	Acc      float32 `json:"acc"`
	MaxCombo int     `json:"combo"`
	Mods     int     `json:"mods"`
	Count300 int     `json:"c300"`
	Count100 int     `json:"c100"`
	Count50  int     `json:"c50"`
	Geki     int     `json:"geki"`
	Katu     int     `json:"katu"`
	Miss     int     `json:"miss"`
	Grade    string  `json:"rank"`
	Status   int     `json:"submission_status"`
	Date     string  `json:"play_date"`
	Perfect  bool    `json:"perfect_score"`
	Map      struct {
		MapID   int    `json:"map_id"`
		Title   string `json:"title"`
		Version string `json:"version"`
		MD5     string `json:"md5"`
		SetID   int    `json:"set_id"`
	} `json:"map"`
}

type AdvancedScore struct {
	ID       int     `json:"id"`
	Score    int     `json:"score"`
	PP       float32 `json:"pp"`
	Acc      float32 `json:"acc"`
	MaxCombo int     `json:"combo"`
	Mods     int     `json:"mods"`
	Count300 int     `json:"c300"`
	Count100 int     `json:"c100"`
	Count50  int     `json:"c50"`
	Geki     int     `json:"geki"`
	Katu     int     `json:"katu"`
	Miss     int     `json:"miss"`
	Grade    string  `json:"rank"`
	Status   int     `json:"submission_status"`
	Date     string  `json:"play_date"`
	Perfect  bool    `json:"perfect_score"`
	Map      struct {
		Title       string  `json:"title"`
		Version     string  `json:"version"`
		SetID       int     `json:"set_id"`
		MapID       int     `json:"map_id"`
		Artist      string  `json:"artist"`
		Creator     string  `json:"creator"`
		LastUpdate  string  `json:"update"`
		TotalLength int     `json:"len"`
		BPM         int     `json:"bpm"`
		CS          int     `json:"cs"`
		AR          int     `json:"ar"`
		OD          int     `json:"od"`
		HP          int     `json:"hp"`
		Diff        float32 `json:"diff"`
		MapStatus   int     `json:"status"`
		MD5         string  `json:"md5"`
		MaxCombo    int     `json:"max_combo"`
	} `json:"map"`
}

// TODO: Disable results for restricted/banned users
// TODO: Treat restricted users as banned
// TODO: Support for rx & ap in GetScores/ScoreInfo

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
			s.status, s.play_time, s.perfect, m.title, m.version,
			m.set_id, m.id
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
			&score.Map.MD5,
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
			&score.Map.Title,
			&score.Map.Version,
			&score.Map.SetID,
			&score.Map.MapID,
		); err != nil {
			log.Println("Error in GetScores")
			return nil, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}

func (db *DB) ScoreInfo(uid int) (AdvancedScore, error) {
	q := `
		SELECT
			s.id, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect,
			m.status, m.id as map_id, m.set_id, m.md5, m.artist, m.title, m.version,
			m.creator, m.last_update, m.total_length, m.max_combo, m.bpm,
			m.cs, m.ar, m.od, m.hp, m.diff
		FROM scores s JOIN maps m ON s.map_md5 = m.md5
		WHERE s.id = ?
	`

	var score AdvancedScore
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, q, uid)

	defer cancel()
	if err := res.Scan(
		&score.ID,
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
		&score.Map.MapStatus,
		&score.Map.MapID,
		&score.Map.SetID,
		&score.Map.MD5,
		&score.Map.Artist,
		&score.Map.Title,
		&score.Map.Version,
		&score.Map.Creator,
		&score.Map.LastUpdate,
		&score.Map.TotalLength,
		&score.Map.MaxCombo,
		&score.Map.BPM,
		&score.Map.CS,
		&score.Map.AR,
		&score.Map.OD,
		&score.Map.HP,
		&score.Map.Diff,
	); err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error in ScoreInfo")
		}

		return score, err
	}

	return score, nil
}
