package parser

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Post struct {
	Text         string
	Size         string
	DateOfBuy    string
	Measurements string
	Model        string
	Picture      string
	Foto         string
}

type Stringer interface {
	String() string
}

func (p *Post) GetMD5Hash() string {
	hash := md5.Sum([]byte(p.Text))
	return hex.EncodeToString(hash[:])
}

func (p Post) String() string {
	return fmt.Sprintf("Text: %s\nSize: %s\nDate: %s\nMeasure: %s\nModel: %s\nPicture: %s\nFoto: %s\n",
		p.Text, p.Size, p.DateOfBuy, p.Measurements, p.Model, p.Picture, p.Foto)
}

func ExampleScrape(st *Post) *Post {
	// Request the HTML page.
	res, err := http.Get("https://groupprice.ru/brands/dstrend/comments")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	firstPost := doc.Find("div.comment-entity").First()
	modelLink := firstPost.Find("a").First().AttrOr("href", "NO LINK")
	if modelLink != "NO LINK" {
		modelLink = "https://groupprice.ru" + modelLink
	}
	// For each item found, get the band and title
	st.Text = firstPost.Find("div.comment-text").Text()
	st.Size = firstPost.Find("div.warehouse-notice").Text()
	st.DateOfBuy = firstPost.Find("div.date").Text()
	st.Measurements = firstPost.Find("div.user-measurements").Text()
	st.Model = modelLink
	st.Picture = firstPost.Find("img").First().AttrOr("src", "NO PICTURE")
	st.Foto = firstPost.Find("a.view").AttrOr("href", "NO FOTO ATTACHED")

	return st
}
