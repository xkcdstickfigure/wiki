package store

import (
	"context"
)

type Article struct {
	Id     string
	SiteId string
	Slug   string
	Title  string
	Source string
}

func (s Store) ArticleGetBySlug(ctx context.Context, siteId string, slug string) (Article, error) {
	var article Article
	err := s.Conn.QueryRow(ctx, "select id, site_id, slug, title, source from article where site_id=$1 and slug=$2", siteId, slug).Scan(&article.Id, &article.SiteId, &article.Slug, &article.Title, &article.Source)
	if err != nil {
		return article, err
	} else {
		return article, nil
	}
}
