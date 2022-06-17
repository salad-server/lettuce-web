package models

import (
	"context"
	"database/sql"
	"log"
	"mini-api/internal/scores"
	"time"
)

type Beatmap struct {
	ID          int     `json:"id"`
	SetID       int     `json:"set_id"`
	Status      string  `json:"status"`
	MD5         string  `json:"md5"`
	Artist      string  `json:"artist"`
	Title       string  `json:"title"`
	Version     string  `json:"version"`
	Creator     string  `json:"creator"`
	LastUpdate  string  `json:"last_update"`
	TotalLength int     `json:"total_length"`
	MaxCombo    int     `json:"max_combo"`
	Plays       int     `json:"plays"`
	Passes      int     `json:"passes"`
	Mode        string  `json:"mode"`
	BPM         float32 `json:"bpm"`
	CS          float32 `json:"cs"`
	AR          float32 `json:"ar"`
	OD          float32 `json:"od"`
	HP          float32 `json:"hp"`
	Difficulty  float32 `json:"diff"`
}

type MapScore struct {
	ID       int     `json:"id"`
	Score    int     `json:"score"`
	Accuracy float32 `json:"acc"`
	MaxCombo int     `json:"max_combo"`
	Mods     int     `json:"mods"`
	Count300 int     `json:"c300"`
	Count100 int     `json:"c100"`
	Count50  int     `json:"c50"`
	Date     string  `json:"play_date"`
	Miss     int     `json:"miss"`
	PP       float32 `json:"pp"`
	User     struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Country  string `json:"country"`
	} `json:"user"`
}

func (db *DB) BeatmapInfo(sid int) (Beatmap, error) {
	q := `
		SELECT
			id, set_id, status, md5, artist, title, version, creator,
			last_update, total_length, max_combo, plays, passes, mode,
			bpm, cs, ar, od, hp, diff
		FROM maps WHERE id = ?
	`

	var bmap Beatmap
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, q, sid)

	defer cancel()
	if err := res.Scan(
		&bmap.ID,
		&bmap.SetID,
		&bmap.Status,
		&bmap.MD5,
		&bmap.Artist,
		&bmap.Title,
		&bmap.Version,
		&bmap.Creator,
		&bmap.LastUpdate,
		&bmap.TotalLength,
		&bmap.MaxCombo,
		&bmap.Plays,
		&bmap.Passes,
		&bmap.Mode,
		&bmap.BPM,
		&bmap.CS,
		&bmap.AR,
		&bmap.OD,
		&bmap.HP,
		&bmap.Difficulty,
	); err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error in BeatmapInfo")
			return bmap, err
		}
	}

	bmap.Status = scores.ConvertStatus(bmap.Status)
	bmap.Mode = scores.ConvertMode(bmap.Mode)
	return bmap, nil
}

func (db *DB) BeatmapSets(bid int) ([]Beatmap, error) {
	q := `
		SELECT
			id, set_id, status, md5, artist, title, version, creator,
			last_update, total_length, max_combo, plays, passes, mode,
			bpm, cs, ar, od, hp, diff
		FROM maps WHERE set_id = (SELECT set_id FROM maps WHERE id = ?)
	`

	var bmaps []Beatmap
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, bid)

	defer cancel()

	if err != nil {
		log.Println("Error in BeatmapSets")
		return nil, err
	}

	for row.Next() {
		var bmap Beatmap

		if err := row.Scan(
			&bmap.ID,
			&bmap.SetID,
			&bmap.Status,
			&bmap.MD5,
			&bmap.Artist,
			&bmap.Title,
			&bmap.Version,
			&bmap.Creator,
			&bmap.LastUpdate,
			&bmap.TotalLength,
			&bmap.MaxCombo,
			&bmap.Plays,
			&bmap.Passes,
			&bmap.Mode,
			&bmap.BPM,
			&bmap.CS,
			&bmap.AR,
			&bmap.OD,
			&bmap.HP,
			&bmap.Difficulty,
		); err != nil {
			log.Println("Error in BeatmapSets")
			return nil, err
		}

		bmap.Status = scores.ConvertStatus(bmap.Status)
		bmap.Mode = scores.ConvertMode(bmap.Mode)
		bmaps = append(bmaps, bmap)
	}

	return bmaps, nil
}

func (db *DB) BeatmapLeaderboard(bid, page int, mode string) ([]MapScore, error) {
	m := scores.Modes[mode]
	q := `
		SELECT
			s.id, s.score, s.acc,
			u.name, u.country, s.userid,
			s.max_combo, s.n300, s.n100, s.n50,
			s.nmiss, s.pp, s.play_time, s.mods
		FROM scores s
		JOIN users u ON u.id = s.userid
		JOIN maps m ON s.map_md5 = m.md5
		WHERE m.id = ? AND s.mode = ? AND s.status = 2
		ORDER BY score DESC LIMIT ?, 10
	`

	var bmap_scores []MapScore
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, bid, m, page*PAGE_LEN)

	defer cancel()

	if err != nil {
		log.Println("Error in BeatmapLeaderboard")
		return nil, err
	}

	for row.Next() {
		var score MapScore

		if err := row.Scan(
			&score.ID,
			&score.Score,
			&score.Accuracy,
			&score.User.Username,
			&score.User.Country,
			&score.User.ID,
			&score.MaxCombo,
			&score.Count300,
			&score.Count100,
			&score.Count50,
			&score.Miss,
			&score.PP,
			&score.Date,
			&score.Mods,
		); err != nil {
			log.Println("Error in BeatmapLeaderboard")
			return nil, err
		}

		bmap_scores = append(bmap_scores, score)
	}

	return bmap_scores, nil
}
