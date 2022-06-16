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
