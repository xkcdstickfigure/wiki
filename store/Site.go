package store

import (
	"context"
)

type Site struct {
	Id           string
	Name         string
	DisplayName  string
	DiscordGuild string
}

func (s Store) SiteGetByName(ctx context.Context, name string) (Site, error) {
	var site Site
	err := s.Conn.QueryRow(ctx, "select id, name, display_name, discord_guild from site where name=$1", name).
		Scan(&site.Id, &site.Name, &site.DisplayName, &site.DiscordGuild)
	return site, err
}
