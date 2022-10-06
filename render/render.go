package render

import (
	"alles/wiki/markup"
)

type PageContext struct {
	Site     string
	PageSlug string
}

func RenderArticle(article markup.Article, pctx PageContext) (string, error) {
	elements, err := renderElements(article.Content.Elements, pctx)
	if err != nil {
		return "", err
	}

	sections, err := renderSections(article.Content.Sections, 0, pctx)
	if err != nil {
		return "", err
	}

	return renderHeader(article.Meta["title"], pctx) + elements + sections, nil
}
