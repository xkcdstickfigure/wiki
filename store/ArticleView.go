package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ArticleView struct {
	Id        string
	SessionId string
	ArticleId string
	Date      time.Time
}

func (s Store) ArticleViewCreate(ctx context.Context, data ArticleView) (ArticleView, error) {
	var view ArticleView
	err := s.Conn.QueryRow(ctx, "insert into article_view (id, session_id, article_id, date) "+
		"values ($1, $2, $3, $4) "+
		"returning id, session_id, article_id, date",
		uuid.New(), data.SessionId, data.ArticleId, time.Now()).
		Scan(&view.Id, &view.SessionId, &view.ArticleId, &view.Date)
	return view, err
}
