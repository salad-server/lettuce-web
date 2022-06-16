package models

import "database/sql"

type DB struct {
	Database *sql.DB
}
