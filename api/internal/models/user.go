package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"mini-api/internal/scores"
)

type Stats struct {
	ID          int     `json:"id"`
	TScore      int64   `json:"total_score"`
	RScore      int64   `json:"ranked_score"`
	PP          int     `json:"pp"`
	Plays       int     `json:"playcount"`
	Acc         float32 `json:"acc"`
	MaxCombo    int     `json:"max_combo"`
	TotalHits   int     `json:"total_hits"`
	ReplayViews int     `json:"watched_replays"`
	XH          int     `json:"count_xh"`
	X           int     `json:"count_x"`
	SH          int     `json:"count_sh"`
	S           int     `json:"count_s"`
	A           int     `json:"count_a"`
}

type Info struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	SafeUsername string `json:"username_safe"`
	Country      string `json:"country"`
	Silence      int    `json:"silence"`
	Register     int    `json:"register_time"`
	LastOnline   int    `json:"last_online"`
	FavMode      string `json:"fav_mode"`
	Playstyle    int    `json:"playstyle"`
	BadgeName    string `json:"badge_name"`
	BadgeIcon    string `json:"badge_icon"`
	Bio          string `json:"bio"`
}

type Session struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Priv     int    `json:"priv"`
}

type Userpage struct {
	Bio       string `json:"bio"`
	Country   string `json:"country"`
	Playstyle int    `json:"playstyle"`
	Email     string `json:"email"`
}

func (db *DB) UserStats(uid int, mode string) (Stats, error) {
	m := scores.Modes[mode]
	q := `
		SELECT
			id, tscore, rscore, pp, plays,
			acc, max_combo, total_hits, replay_views,
			xh_count, x_count, sh_count, s_count, a_count
		FROM stats
		WHERE id = ? AND mode = ?
	`

	var stats Stats
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, q, uid, m)

	defer cancel()
	if err := res.Scan(
		&stats.ID,
		&stats.TScore,
		&stats.RScore,
		&stats.PP,
		&stats.Plays,
		&stats.Acc,
		&stats.MaxCombo,
		&stats.TotalHits,
		&stats.ReplayViews,
		&stats.XH,
		&stats.X,
		&stats.SH,
		&stats.S,
		&stats.A,
	); err != nil && err != sql.ErrNoRows {
		log.Println("Error in UserStats")
		return stats, err
	}

	return stats, nil
}

func (db *DB) UserInfo(uid int) (Info, error) {
	var info Info
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, `
		SELECT
			id, name, safe_name, country, silence_end,
			creation_time, latest_activity, preferred_mode,
			play_style, COALESCE(custom_badge_name, ''), COALESCE(custom_badge_icon, ''),
			COALESCE(userpage_content, '')
		FROM users WHERE id = ?
	`, uid)

	defer cancel()
	if err := res.Scan(
		&info.ID,
		&info.Username,
		&info.SafeUsername,
		&info.Country,
		&info.Silence,
		&info.Register,
		&info.LastOnline,
		&info.FavMode,
		&info.Playstyle,
		&info.BadgeName,
		&info.BadgeIcon,
		&info.Bio,
	); err != nil && err != sql.ErrNoRows {
		log.Println("Error in UserInfo")
		return info, err
	}

	info.FavMode = scores.ConvertMode(info.FavMode)
	return info, nil
}

func (db *DB) UserCreds(email string) (string, error) {
	var pw string
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, "SELECT pw_bcrypt FROM users WHERE email = ?", email)

	defer cancel()

	if err := res.Scan(&pw); err != nil {
		log.Println("Error in UserCreds")
		return "", err
	}

	return pw, nil
}

func (db *DB) UserSession(email string) (Session, error) {
	var session Session
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, "SELECT id, name, priv FROM users WHERE email = ?", email)

	defer cancel()

	if err := res.Scan(
		&session.ID,
		&session.Username,
		&session.Priv,
	); err != nil {
		log.Println("Error in UserSession")
		return session, err
	}

	return session, nil
}

func (db *DB) ProfileEdit(uid int) (Userpage, error) {
	var page Userpage
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res := db.Database.QueryRowContext(c, "SELECT COALESCE(userpage_content, ''), country, play_style, email FROM users WHERE id = ?", uid)

	defer cancel()

	if err := res.Scan(
		&page.Bio,
		&page.Country,
		&page.Playstyle,
		&page.Email,
	); err != nil {
		log.Println("Error in ProfileEdit")
		return page, err
	}

	return page, nil
}

func (db *DB) ProfileUpdate(bio, country, email string, playstyle, uid int) error {
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := db.Database.ExecContext(c, `
		UPDATE users SET
			userpage_content = ?,
			country = ?,
			email = ?,
			play_style = ?
		WHERE id = ?
	`, bio, country, email, playstyle, uid)

	defer cancel()

	if err != nil {
		log.Println("Error in ProfileUpdate")
		return err
	}

	return nil
}

func (db *DB) LastSeen(uid int) error {
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := db.Database.ExecContext(c, "UPDATE users SET latest_activity = UNIX_TIMESTAMP() WHERE id = ?", uid)

	defer cancel()

	if err != nil {
		log.Println("Error in LastSeen")
		return err
	}

	return nil
}
