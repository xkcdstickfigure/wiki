package store

import (
	"github.com/jackc/pgx/v5"
)

type Store struct {
	Conn *pgx.Conn
}
