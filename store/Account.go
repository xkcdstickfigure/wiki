package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id            string
	GoogleId      string
	DiscordId     sql.NullString
	Name          string
	Email         string
	EmailVerified bool
	Avatar        string
	CreatedAt     time.Time
}

func (s Store) AccountCreate(ctx context.Context, data Account) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(
		ctx,
		"insert into account "+
			"(id, google_id, name, email, email_verified, avatar, created_at) "+
			"values ($1, $2, $3, $4, $5, $6, $7) "+
			"on conflict (google_id) do update set name=$3, email=$4, email_verified=$5, avatar=$6"+
			"returning id, google_id, discord_id, name, email, email_verified, avatar, created_at",
		uuid.New(), data.GoogleId, data.Name, data.Email, data.EmailVerified, data.Avatar, time.Now()).
		Scan(&account.Id, &account.GoogleId, &account.DiscordId, &account.Name, &account.Email, &account.EmailVerified, &account.Avatar, &account.CreatedAt)
	return account, err
}

func (s Store) AccountSetDiscord(ctx context.Context, id string, discordId string) error {
	_, err := s.Conn.Exec(ctx, "update account set discord_id=$2 where id=$1", id, discordId)
	if err != nil {
		return err
	}

	_, err = s.Conn.Exec(ctx, "update session set discord_id=$2 where account_id=$1", id, discordId)
	return err
}
