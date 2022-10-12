package store

import (
	"context"
	"time"
)

type DiscordUser struct {
	Id            string
	Username      string
	Discriminator string
	Avatar        string
	MfaEnabled    bool
	Locale        string
	Email         string
	EmailVerified bool
	AccessToken   string
	RefreshToken  string
	FirstAuthAt   time.Time
	LastAuthAt    time.Time
}

func (s Store) DiscordUserCreate(ctx context.Context, user DiscordUser) (DiscordUser, error) {
	var u DiscordUser
	err := s.Conn.QueryRow(
		ctx,
		"insert into discord_user "+
			"(id, username, discriminator, avatar, mfa_enabled, locale, email, email_verified, access_token, refresh_token, first_auth_at, last_auth_at) "+
			"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $11) "+
			"on conflict (id) do update set username=$2, discriminator=$3, avatar=$4, mfa_enabled=$5, locale=$6, email=$7, email_verified=$8, access_token=$9, refresh_token=$10, last_auth_at=$11 "+
			"returning id, username, discriminator, avatar, mfa_enabled, locale, email, email_verified, access_token, refresh_token, first_auth_at, last_auth_at",
		user.Id, user.Username, user.Discriminator, user.Avatar, user.MfaEnabled, user.Locale, user.Email, user.EmailVerified, user.AccessToken, user.RefreshToken, time.Now()).
		Scan(&u.Id, &u.Username, &u.Discriminator, &u.Avatar, &u.MfaEnabled, &u.Locale, &u.Email, &u.EmailVerified, &u.AccessToken, &u.RefreshToken, &u.FirstAuthAt, &u.LastAuthAt)
	return u, err
}
