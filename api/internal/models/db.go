package models

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type DB struct {
	Database *sql.DB
	Client *redis.Client
}
