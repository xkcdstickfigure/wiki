package render

import (
	"alles/wiki/markup"
)

type PageContext struct {
	Site     string
	PageSlug string
}

func RenderArticle(article markup.Article, pctx PageContext) (string, error) {
	element, err := renderElements(article.Content.Elements, pctx)
	if err != nil {
		return "", err
	}

	return renderHeader(article.Meta["title"], pctx) + element, nil
}
