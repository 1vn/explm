package main

import (
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

	memeName := os.Args[1]

	url := fmt.Sprintf("http://knowyourmeme.com/memes/%s", memeName)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln("Error, unable to parse webpage: ", err)
	}

	log.Println(doc.Find("h1").First().Find("a").Text())
}
