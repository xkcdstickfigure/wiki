package markup

type Article struct {
	Meta    Meta
	Infobox Infobox
	Content Content
}

func ParseArticle(str string) (Article, error) {
	parts, err := splitParts(str)
	if err != nil {
		return Article{}, err
	}

	meta, err := parseMeta(parts["meta"])
	if err != nil {
		return Article{}, err
	}

	infobox, err := parseInfobox(parts["infobox"])
	if err != nil {
		return Article{}, err
	}

	content, err := parseContent(parts["content"])
	if err != nil {
		return Article{}, err
	}

	return Article{
		Meta:    meta,
		Infobox: infobox,
		Content: content,
	}, nil
}
