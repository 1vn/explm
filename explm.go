package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s {$meme_name}\n", os.Args[0])
		os.Exit(1)
	}

	meme_name := os.Args[1]

	var url bytes.Buffer
	url.WriteString("http://knowyourmeme.com/memes/")
	url.WriteString(meme_name)

	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		log.Fatalln("Error, unable to parse webpage: ", err)
	}

	log.Println(doc.Find("h1").First().Find("a").Text())
}
