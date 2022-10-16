package store

import (
	"context"
	"database/sql"
	"time"

	"alles/wiki/random"

	"github.com/google/uuid"
)

type Account struct {
	Id            string
	Number        int
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
			"(id, number, google_id, name, email, email_verified, avatar, created_at) "+
			"values ($1, $2, $3, $4, $5, $6, $7, $8) "+
			"on conflict (google_id) do update set name=$4, email=$5, email_verified=$6, avatar=$7"+
			"returning id, number, google_id, discord_id, name, email, email_verified, avatar, created_at",
		uuid.New(), random.Number(10), data.GoogleId, data.Name, data.Email, data.EmailVerified, data.Avatar, time.Now()).
		Scan(&account.Id, &account.Number, &account.GoogleId, &account.DiscordId, &account.Name, &account.Email, &account.EmailVerified, &account.Avatar, &account.CreatedAt)
	return account, err
}

func (s Store) AccountGet(ctx context.Context, id string) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(ctx, "select id, number, google_id, discord_id, name, email, email_verified, avatar, created_at from account where id=$1", id).
		Scan(&account.Id, &account.Number, &account.GoogleId, &account.DiscordId, &account.Name, &account.Email, &account.EmailVerified, &account.Avatar, &account.CreatedAt)
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
