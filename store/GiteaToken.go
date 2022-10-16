package store

import (
	"context"
	"time"

	"alles/wiki/randstr"

	"github.com/google/uuid"
)

type GiteaToken struct {
	Id        string
	AccountId string
	Token     string
	CreatedAt time.Time
}

func (s Store) GiteaTokenCreate(ctx context.Context, accountId string) (GiteaToken, error) {
	var token GiteaToken
	err := s.Conn.QueryRow(ctx, "insert into gitea_token (id, account_id, token, created_at) "+
		"values ($1, $2, $3, $4) "+
		"returning id, account_id, token, created_at",
		uuid.New(), accountId, randstr.Generate(32), time.Now()).
		Scan(&token.Id, &token.AccountId, &token.Token, &token.CreatedAt)
	return token, err
}

func (s Store) GiteaTokenUse(ctx context.Context, c string) (GiteaToken, error) {
	var token GiteaToken
	err := s.Conn.QueryRow(ctx, "select id, account_id, token, created_at from gitea_token where token=$1", c).
		Scan(&token.Id, &token.AccountId, &token.Token, &token.CreatedAt)

	if err != nil {
		return token, err
	}

	_, err = s.Conn.Exec(ctx, "delete from gitea_token where token=$1", c)

	return token, err
}
