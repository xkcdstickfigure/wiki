package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"alles/wiki/markup"
)

//go:embed eye_of_cthulhu.txt
var source string

func main() {
	article, err := markup.ParseArticle(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := json.MarshalIndent(article, "", "	")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
