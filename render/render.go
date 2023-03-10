package render

import (
	"alles/wiki/markup"
)

type PageContext struct {
	Title         string
	Site          string
	Domain        string
	Slug          string
	StorageOrigin string
}

func RenderArticle(article markup.Article, pctx PageContext) (string, error) {
	// infobox
	infobox, err := renderInfobox(article, pctx)
	if err != nil {
		return "", err
	}

	// header
	header := renderHeader(pctx)

	// elements
	elements, err := renderElements(article.Content.Elements, pctx)
	if err != nil {
		return "", err
	}

	// table of contents
	toc := renderToc(article.Content.Sections)

	// sections
	sections, err := renderSections(article.Content.Sections, 0, pctx)
	if err != nil {
		return "", err
	}

	// return
	return `<section>` + infobox + header + elements + toc + `</section>` + sections, nil
}
