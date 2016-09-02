package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s {$meme_name}\n", os.Args[0])
		os.Exit(1)
	}

	get_meme(os.Args[1])
}

func get_meme(meme_name string) {
	var url bytes.Buffer
	url.WriteString("http://knowyourmeme.com/memes/")
	url.WriteString(meme_name)

	doc, err := goquery.NewDocument(url.String())
	if err != nil {
		log.Fatalln("Error, unable to parse webpage: ", err)
	}

	log.Println(doc.Find("h1").First().Find("a").Text())
}

// http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html

func parse_html(body io.Reader) {
	ttz := html.NewTokenizer(body)
	for {
		tt := ttz.Next()
		if tt == html.ErrorToken {
			break
		}
		// consume the token
	}
}
