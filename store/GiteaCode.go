package store

import (
	"context"
	"time"

	"alles/wiki/random"

	"github.com/google/uuid"
)

type GiteaCode struct {
	Id        string
	AccountId string
	Code      string
	CreatedAt time.Time
}

func (s Store) GiteaCodeCreate(ctx context.Context, accountId string) (GiteaCode, error) {
	var code GiteaCode
	err := s.Conn.QueryRow(ctx, "insert into gitea_code (id, account_id, code, created_at) "+
		"values ($1, $2, $3, $4) "+
		"returning id, account_id, code, created_at",
		uuid.New(), accountId, random.String(32), time.Now()).
		Scan(&code.Id, &code.AccountId, &code.Code, &code.CreatedAt)
	return code, err
}

func (s Store) GiteaCodeUse(ctx context.Context, c string) (GiteaCode, error) {
	var code GiteaCode
	err := s.Conn.QueryRow(ctx, "select id, account_id, code, created_at from gitea_code where code=$1", c).
		Scan(&code.Id, &code.AccountId, &code.Code, &code.CreatedAt)

	if err != nil {
		return code, err
	}

	_, err = s.Conn.Exec(ctx, "delete from gitea_code where code=$1", c)

	return code, err
}
