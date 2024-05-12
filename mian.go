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

	_, err := getAllSiteMapUrls(c, sitemapLink)
	
	if err != nil {
		log.Fatalf("Error getting links from sitemap: %v\n", err)
	}

}

func getAllSiteMapUrls(c *colly.Collector, sitemapLink string) ([]string, error) {
	urls := []string{}

	c.OnHTML("#sitemap-app-list-wrapper a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("link:", link)
		urls = append(urls, link)
	})

	err := c.Visit(sitemapLink)
	if err != nil {
		return nil, err
	}

	return urls, nil

}
