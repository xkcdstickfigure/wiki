package store

import (
	"context"
)

type DiscordGuild struct {
	Id   string
	Name string
	Icon string
}

func (s Store) DiscordGuildCreate(ctx context.Context, guild DiscordGuild, userId string) error {
	_, err := s.Conn.Exec(ctx, "insert into discord_guild (id, name, icon) values ($1, $2, $3) on conflict (id) do update set name=$2, icon=$3", guild.Id, guild.Name, guild.Icon)
	if err != nil {
		return err
	}

	_, err = s.Conn.Exec(ctx, "insert into discord_member (user_id, guild_id) values ($1, $2) on conflict do nothing", userId, guild.Id)
	return err
}
