package store

import (
	"context"
	"time"

	"alles/wiki/randstr"

	"github.com/google/uuid"
)

type DiscordState struct {
	Id        string
	SessionId string
	Token     string
	Site      string
	CreatedAt time.Time
}

func (s Store) DiscordStateCreate(ctx context.Context, data DiscordState) (DiscordState, error) {
	var state DiscordState
	err := s.Conn.QueryRow(ctx, "insert into discord_state (id, session_id, token, site, created_at) "+
		"values ($1, $2, $3, $4, $5) "+
		"returning id, session_id, token, site, created_at",
		uuid.New(), data.SessionId, randstr.Generate(32), data.Site, time.Now()).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Site, &state.CreatedAt)
	return state, err
}

func (s Store) DiscordStateUse(ctx context.Context, token string) (DiscordState, error) {
	var state DiscordState
	err := s.Conn.QueryRow(ctx, "select id, session_id, token, site, created_at from discord_state where token=$1", token).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Site, &state.CreatedAt)

	if err != nil {
		return state, err
	}

	_, err = s.Conn.Exec(ctx, "delete from discord_state where token=$1", token)

	return state, err
}
