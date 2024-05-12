package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {

	sitemapLink := "https://ridge.com/tools/sitemap"

	c := colly.NewCollector(
		colly.AllowedDomains("ridge.com"), // Allowed domain, without protocol
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("#sitemap-app-list-wrapper a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("link:", link)
	})

	// Start the collector
	c.Visit(sitemapLink)

}
