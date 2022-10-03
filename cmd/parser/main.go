package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"alles/wiki/markup"
)

//go:embed eye_of_cthulhu.txt
var article string

func main() {
	parts, err := markup.SplitParts(article)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	meta, err := markup.ParseMeta(parts["meta"])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := json.MarshalIndent(meta, "", "	")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
