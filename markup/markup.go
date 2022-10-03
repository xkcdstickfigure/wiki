package markup

import "errors"

type Article struct {
	Meta    Meta
	Infobox Infobox
	Content Content
}

func ParseArticle(str string) (Article, error) {
	parts, err := SplitParts(str)
	if err != nil {
		return Article{}, err
	}

	if len(parts["meta"]) == 0 {
		return Article{}, errors.New("meta part is required")
	}

	if len(parts["infobox"]) == 0 {
		return Article{}, errors.New("infobox part is required")
	}

	if len(parts["content"]) == 0 {
		return Article{}, errors.New("content part is required")
	}

	meta, err := ParseMeta(parts["meta"])
	if err != nil {
		return Article{}, err
	}

	infobox, err := ParseInfobox(parts["infobox"])
	if err != nil {
		return Article{}, err
	}

	content, err := ParseContent(parts["content"])
	if err != nil {
		return Article{}, err
	}

	return Article{
		Meta:    meta,
		Infobox: infobox,
		Content: content,
	}, nil
}
