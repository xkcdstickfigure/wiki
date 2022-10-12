package store

import (
	"context"
	"database/sql"
)

type Site struct {
	Id            string
	Name          string
	DisplayName   string
	DiscordServer sql.NullString
}

func (s Store) SiteGetByName(ctx context.Context, name string) (Site, error) {
	var site Site
	err := s.Conn.QueryRow(ctx, "select id, name, display_name, discord_server from site where name=$1", name).Scan(&site.Id, &site.Name, &site.DisplayName, &site.DiscordServer)
	if err != nil {
		return site, err
	} else {
		return site, nil
	}
}
