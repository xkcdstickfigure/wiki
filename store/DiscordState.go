package store

import (
	"context"
	"time"

	"alles/wiki/random"

	"github.com/google/uuid"
)

type DiscordState struct {
	Id        string
	SessionId string
	Token     string
	Action    string
	Value     string
	CreatedAt time.Time
}

func (s Store) DiscordStateCreate(ctx context.Context, data DiscordState) (DiscordState, error) {
	var state DiscordState
	err := s.Conn.QueryRow(ctx, "insert into discord_state (id, session_id, token, action, value, created_at) "+
		"values ($1, $2, $3, $4, $5, $6) "+
		"returning id, session_id, token, action, value, created_at",
		uuid.New(), data.SessionId, random.String(32), data.Action, data.Value, time.Now()).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Action, &state.Value, &state.CreatedAt)
	return state, err
}

func (s Store) DiscordStateUse(ctx context.Context, token string) (DiscordState, error) {
	var state DiscordState
	err := s.Conn.QueryRow(ctx, "select id, session_id, token, action, value, created_at from discord_state where token=$1", token).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Action, &state.Value, &state.CreatedAt)

	if err != nil {
		return state, err
	}

	_, err = s.Conn.Exec(ctx, "delete from discord_state where token=$1", token)

	return state, err
}
