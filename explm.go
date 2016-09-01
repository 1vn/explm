package main

import (
    "bytes"
    "fmt"
    "golang.org/x/net/html"
    "io"
    "log"
    "net/http"
    "os"
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

    rs, err := http.Get(url.String())
    if err != nil {
        log.Fatal(err)
    } else {
        defer rs.Body.Close()
        if err != nil {
            log.Fatal(err)
        } else {
            parse_html(rs.Body)
        }
    }
}

// http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html

func parse_html (body io.Reader) {
    ttz := html.NewTokenizer(body)
    for {
        tt := ttz.Next()
        if tt == html.ErrorToken{
            break
        }
        // consume the token
    }
}