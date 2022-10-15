package store

import (
	"context"
	"time"

	"alles/wiki/randstr"

	"github.com/google/uuid"
)

type AuthState struct {
	Id        string
	SessionId string
	Token     string
	Redirect  string
	CreatedAt time.Time
}

func (s Store) AuthStateCreate(ctx context.Context, data AuthState) (AuthState, error) {
	var state AuthState
	err := s.Conn.QueryRow(ctx, "insert into auth_state (id, session_id, token, redirect, created_at) "+
		"values ($1, $2, $3, $4, $5) "+
		"returning id, session_id, token, redirect, created_at",
		uuid.New(), data.SessionId, randstr.Generate(32), data.Redirect, time.Now()).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Redirect, &state.CreatedAt)
	return state, err
}

func (s Store) AuthStateUse(ctx context.Context, token string) (AuthState, error) {
	var state AuthState
	err := s.Conn.QueryRow(ctx, "select id, session_id, token, redirect, created_at from auth_state where token=$1", token).
		Scan(&state.Id, &state.SessionId, &state.Token, &state.Redirect, &state.CreatedAt)

	if err != nil {
		return state, err
	}

	_, err = s.Conn.Exec(ctx, "delete from auth_state where token=$1", token)

	return state, err
}
