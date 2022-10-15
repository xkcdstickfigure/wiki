package store

import (
	"context"
	"database/sql"
	"time"

	"alles/wiki/randstr"

	"github.com/google/uuid"
)

type Session struct {
	Id        string
	Token     string
	Address   string
	UserAgent string
	AccountId sql.NullString
	DiscordId sql.NullString
	CreatedAt time.Time
}

func (s Store) SessionGetByToken(ctx context.Context, token string) (Session, error) {
	var session Session
	err := s.Conn.QueryRow(ctx, "select id, token, address, user_agent, account_id, discord_id, created_at from session where token=$1", token).
		Scan(&session.Id, &session.Token, &session.Address, &session.UserAgent, &session.AccountId, &session.DiscordId, &session.CreatedAt)
	return session, err
}

func (s Store) SessionCreate(ctx context.Context, data Session) (Session, error) {
	var session Session
	err := s.Conn.QueryRow(ctx, "insert into session (id, token, address, user_agent, created_at) "+
		"values ($1, $2, $3, $4, $5) "+
		"returning id, token, address, user_agent, created_at",
		uuid.New(), randstr.Generate(32), data.Address, data.UserAgent, time.Now()).
		Scan(&session.Id, &session.Token, &session.Address, &session.UserAgent, &session.CreatedAt)
	return session, err
}

func (s Store) SessionSetAccount(ctx context.Context, id string, accountId string) error {
	_, err := s.Conn.Exec(ctx, "update session set account_id=$2 where id=$1", id, accountId)
	return err
}

func (s Store) SessionSetDiscord(ctx context.Context, id string, discordId string) error {
	_, err := s.Conn.Exec(ctx, "update session set discord_id=$2 where id=$1", id, discordId)
	return err
}
