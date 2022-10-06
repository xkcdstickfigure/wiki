package markup

import "strings"

type Article struct {
	Meta    Meta
	Image   Image
	Infobox Infobox
	Content Content
}

func ParseArticle(str string) (Article, error) {
	parts, err := splitParts(str)
	if err != nil {
		return Article{}, err
	}

	// meta
	meta, err := parseMeta(parts["meta"])
	if err != nil {
		return Article{}, err
	}

	// image
	imageSource := strings.Split(meta["image"], " ")[0]
	imageText, err := parseText(strings.TrimPrefix(meta["image"], imageSource))
	if err != nil {
		return Article{}, err
	}

	// infobox
	infobox, err := parseInfobox(parts["infobox"])
	if err != nil {
		return Article{}, err
	}

	// content
	content, err := parseContent(parts["content"])
	if err != nil {
		return Article{}, err
	}

	// return
	return Article{
		Meta: meta,
		Image: Image{
			Source: imageSource,
			Text:   imageText,
		},
		Infobox: infobox,
		Content: content,
	}, nil
}
