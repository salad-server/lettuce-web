package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mini-api/internal/scores"
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
		MapID     int    `json:"map_id"`
		Title     string `json:"title"`
		Version   string `json:"version"`
		MD5       string `json:"md5"`
		SetID     int    `json:"set_id"`
		MapStatus string `json:"status"`
	} `json:"map"`
}

type AdvancedScore struct {
	ID       int     `json:"id"`
	UID      int     `json:"uid"`
	Username string  `json:"username"`
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
	Mode     string  `json:"play_mode"`
	Map      struct {
		Title       string  `json:"title"`
		Version     string  `json:"version"`
		SetID       int     `json:"set_id"`
		MapID       int     `json:"map_id"`
		Artist      string  `json:"artist"`
		Creator     string  `json:"creator"`
		LastUpdate  string  `json:"update"`
		TotalLength int     `json:"len"`
		BPM         float32 `json:"bpm"`
		CS          float32 `json:"cs"`
		AR          float32 `json:"ar"`
		OD          float32 `json:"od"`
		HP          float32 `json:"hp"`
		Diff        float32 `json:"diff"`
		MapStatus   string  `json:"status"`
		MD5         string  `json:"md5"`
		MaxCombo    int     `json:"max_combo"`
	} `json:"map"`
}

type ScoreInfoResponse struct {
	AdvancedScore
	Pinned bool `json:"pinned"`
}

type Records struct {
	Score AdvancedScore `json:"score"`
	PP    AdvancedScore `json:"pp"`
}

func (db *DB) GetScores(best bool, uid, page int, mode string) ([]Score, error) {
	m := scores.Modes[mode]
	sort := "s.id"
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
			m.set_id, m.id, m.status
		FROM scores s
		INNER JOIN maps m ON s.map_md5 = m.md5
		INNER JOIN users u ON s.userid = u.id
		WHERE s.userid = ? AND s.mode = ? AND u.priv & 1 %s
		ORDER BY %s DESC LIMIT ?, 10
	`, ranked, sort)

	var all_scores []Score
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, uid, m, page*PAGE_LEN)

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
			&score.Map.MapStatus,
		); err != nil {
			log.Println("Error in GetScores")
			return nil, err
		}

		score.Map.MapStatus = scores.ConvertStatus(score.Map.MapStatus)
		all_scores = append(all_scores, score)
	}

	return all_scores, nil
}

func (db *DB) ScoreInfo(uid int) (ScoreInfoResponse, error) {
	q := `
		SELECT
			s.id, s.userid, u.name, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect, s.mode,
			m.status, m.id AS map_id, m.set_id, m.md5, m.artist, m.title, m.version,
			m.creator, m.last_update, m.total_length, m.max_combo, m.bpm,
			m.cs, m.ar, m.od, m.hp, m.diff
		FROM scores s
		JOIN maps m ON s.map_md5 = m.md5
		JOIN users u ON s.userid = u.id
		WHERE s.id = ? AND u.priv & 1
	`

	var score ScoreInfoResponse
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, q, uid)

	defer cancel()
	if err := res.Scan(
		&score.ID,
		&score.UID,
		&score.Username,
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
		&score.Mode,
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
	); err != nil && err != sql.ErrNoRows {
		log.Println("Error in ScoreInfo")
		return score, err
	}

	score.Map.MapStatus = scores.ConvertStatus(score.Map.MapStatus)
	score.Mode = scores.ConvertMode(score.Mode)
	score.Pinned = db.PinnedExists(score.ID)

	return score, nil
}

func (db *DB) Records() (map[string]Records, error) {
	cache, err := db.Client.Get(context.Background(), "mini:records").Result()
	records := map[string]Records{
		"vn!std":   {},
		"vn!taiko": {},
		"vn!catch": {},
		"vn!mania": {},
		"rx!std":   {},
		"rx!taiko": {},
		"rx!catch": {},
		"ap!std":   {},
	}

	// load from cache mini:records
	if err == nil {
		if err := json.Unmarshal([]byte(cache), &records); err != nil {
			log.Println("cache found but could not unmarshal", err)
		} else {
			return records, nil
		}
	}

	tmpl := `
		SELECT
			s.id, s.userid, u.name, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect, s.mode,
			m.status, m.id AS map_id, m.set_id, m.md5, m.artist, m.title, m.version,
			m.creator, m.last_update, m.total_length, m.max_combo, m.bpm,
			m.cs, m.ar, m.od, m.hp, m.diff
		FROM scores s
		JOIN maps m ON s.map_md5 = m.md5
		JOIN users u ON s.userid = u.id
		WHERE s.mode = ? AND m.status IN (2, 3, 4) AND s.status = 2 AND u.priv & 1
		ORDER BY %s DESC LIMIT 1
	`

	for k := range records {
		var record_pp AdvancedScore
		var record_score AdvancedScore
		pp := fmt.Sprintf(tmpl, "s.pp")
		score := fmt.Sprintf(tmpl, "s.score")

		v := scores.Modes[k]

		cpp, cancel_pp := context.WithTimeout(context.Background(), 3*time.Second)
		cscore, cancel_score := context.WithTimeout(context.Background(), 3*time.Second)
		res_pp := db.Database.QueryRowContext(cpp, pp, v)
		res_score := db.Database.QueryRowContext(cscore, score, v)

		defer cancel_pp()
		defer cancel_score()

		if err := res_pp.Scan(
			&record_pp.ID, &record_pp.UID,
			&record_pp.Username, &record_pp.Score,
			&record_pp.PP, &record_pp.Acc,
			&record_pp.MaxCombo, &record_pp.Mods,
			&record_pp.Count300, &record_pp.Count100,
			&record_pp.Count50, &record_pp.Miss, &record_pp.Geki,
			&record_pp.Katu, &record_pp.Grade,
			&record_pp.Status, &record_pp.Date,
			&record_pp.Perfect, &record_pp.Mode,
			&record_pp.Map.MapStatus, &record_pp.Map.MapID,
			&record_pp.Map.SetID, &record_pp.Map.MD5,
			&record_pp.Map.Artist, &record_pp.Map.Title,
			&record_pp.Map.Version, &record_pp.Map.Creator,
			&record_pp.Map.LastUpdate, &record_pp.Map.TotalLength,
			&record_pp.Map.MaxCombo, &record_pp.Map.BPM,
			&record_pp.Map.CS, &record_pp.Map.AR,
			&record_pp.Map.OD, &record_pp.Map.HP,
			&record_pp.Map.Diff,
		); err != nil && err != sql.ErrNoRows {
			log.Println("Error in Records")
			return records, err
		}

		record_pp.Map.MapStatus = scores.ConvertStatus(record_pp.Map.MapStatus)
		record_pp.Mode = scores.ConvertMode(record_pp.Mode)

		if err := res_score.Scan(
			&record_score.ID, &record_score.UID,
			&record_score.Username, &record_score.Score,
			&record_score.PP, &record_score.Acc,
			&record_score.MaxCombo, &record_score.Mods,
			&record_score.Count300, &record_score.Count100,
			&record_score.Count50, &record_score.Miss, &record_score.Geki,
			&record_score.Katu, &record_score.Grade,
			&record_score.Status, &record_score.Date,
			&record_score.Perfect, &record_score.Mode,
			&record_score.Map.MapStatus, &record_score.Map.MapID,
			&record_score.Map.SetID, &record_score.Map.MD5,
			&record_score.Map.Artist, &record_score.Map.Title,
			&record_score.Map.Version, &record_score.Map.Creator,
			&record_score.Map.LastUpdate, &record_score.Map.TotalLength,
			&record_score.Map.MaxCombo, &record_score.Map.BPM,
			&record_score.Map.CS, &record_score.Map.AR,
			&record_score.Map.OD, &record_score.Map.HP,
			&record_score.Map.Diff,
		); err != nil && err != sql.ErrNoRows {
			log.Println("Error in Records")
			return records, err
		}

		record_score.Map.MapStatus = scores.ConvertStatus(record_score.Map.MapStatus)
		record_score.Mode = scores.ConvertMode(record_score.Mode)

		records[k] = Records{
			PP:    record_pp,
			Score: record_score,
		}
	}

	// 4 hours
	js, jerr := json.Marshal(records)
	if err := db.Client.Set(context.Background(), "mini:records", string(js), time.Hour*4).Err(); jerr != nil || err != nil {
		log.Println("Could not set mini:records.", err, jerr)
	}

	return records, nil
}

func (db *DB) GetPinned(uid, page int, mode string) ([]Score, error) {
	m := scores.Modes[mode]
	q := `
		SELECT
			s.id, s.map_md5, s.score, s.pp, s.acc, s.max_combo, s.mods,
			s.n300, s.n100, s.n50, s.nmiss, s.ngeki, s.nkatu, s.grade, 
			s.status, s.play_time, s.perfect, m.title, m.version,
			m.set_id, m.id, m.status
		FROM pinned p
		JOIN scores s ON s.id = p.id
		JOIN maps m ON s.map_md5 = m.md5
		JOIN users u ON s.userid = u.id
		WHERE s.userid = ? AND s.mode = ? AND u.priv & 1
		ORDER BY p.pintime DESC
		LIMIT ?, 5
	`

	var pinned []Score
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	row, err := db.Database.QueryContext(c, q, uid, m, page*(PAGE_LEN/2))

	defer cancel()

	if err != nil {
		log.Println("Error in GetPinned")
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
			log.Println("Error in GetPinned")
			return nil, err
		}

		score.Map.MapStatus = scores.ConvertStatus(score.Map.MapStatus)
		pinned = append(pinned, score)
	}

	return pinned, nil
}

func (db *DB) ScoreExists(uid, sid int) bool {
	var score bool
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, `
		SELECT 1 FROM scores s
		JOIN users u ON s.userid = u.id
		WHERE s.id = ? AND s.userid = ? AND u.priv & 1 AND s.status != 0
	`, sid, uid)

	defer cancel()

	if err := res.Scan(&score); err != nil && err != sql.ErrNoRows {
		log.Println("Error in ScoreExists")
		return false
	}

	return score
}

func (db *DB) PinnedExists(id int) bool {
	var pinned bool
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, "SELECT 1 FROM pinned WHERE id = ?", id)

	defer cancel()

	if err := res.Scan(&pinned); err != nil && err != sql.ErrNoRows {
		log.Println("Error in PinnedExists")
		return false
	}

	return pinned
}
