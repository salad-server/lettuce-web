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

func (db *DB) BeatmapLeaderboard(bid, page int, mode string) ([]Score, error) {
	m := scores.Modes[mode]
	q := `
		SELECT
			s.id, s.map_md5, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect, m.title, m.version,
			m.set_id, m.id, m.status
		FROM scores s
		INNER JOIN maps m ON s.map_md5 = m.md5
		WHERE m.id = ? AND s.mode = ? AND s.status = 2
		ORDER BY pp DESC LIMIT ?, 10
	`

	var bmap_scores []Score
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, bid, m, page*PAGE_LEN)

	defer cancel()

	if err != nil {
		log.Println("Error in BeatmapLeaderboard")
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
			&score.Map.MapStatus,
		); err != nil {
			log.Println("Error in BeatmapLeaderboard")
			return nil, err
		}

		score.Map.MapStatus = scores.ConvertStatus(score.Map.MapStatus)
		bmap_scores = append(bmap_scores, score)
	}

	return bmap_scores, nil
}
