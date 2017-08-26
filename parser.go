package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func parseTabPage(doc *goquery.Document) {

	// ultimate-guitar
	tabContent := doc.Find("pre.js-tab-content").Text()
	title := doc.Find(".t_title h1").Text()
	artist := doc.Find(".t_title .t_autor a").Text()
	fmt.Println(title)
	fmt.Println(artist)
	fmt.Println(tabContent)
}