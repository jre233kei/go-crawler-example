package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://golang.org/dl/")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	version := doc.Find("div", "class", "toggleVisible").Attrs()["id"]
	fmt.Println(version)
}
