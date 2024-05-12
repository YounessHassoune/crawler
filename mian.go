package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	sitemapLink := "https://ridge.com/tools/sitemap"

	c := colly.NewCollector(
		colly.AllowedDomains("ridge.com"), // Allowed domain, without protocol
	)

	urls, err := getAllSiteMapUrls(c, sitemapLink)

	if err != nil {
		log.Fatalf("Error getting links from sitemap: %v\n", err)
	}

	fmt.Println("Number of URLs collected:", len(urls))

}

func getAllSiteMapUrls(c *colly.Collector, sitemapLink string) ([]string, error) {
	urls := []string{}

	c.OnHTML("#sitemap-app-list-wrapper a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		u, err := url.Parse(link)

		if err != nil {
			return
		}

		if page := u.Query().Get("page"); page != "" {
			err := c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
			if err != nil {
				log.Printf("Error visiting pagination link %s: %v\n", e.Attr("href"), err)
			}
		}

		fmt.Println("link:", link)

		if !strings.HasPrefix(link, "/tools/sitemap?") {
			urls = append(urls, link)
		}
	})

	err := c.Visit(sitemapLink)

	if err != nil {
		return nil, err
	}

	return urls, nil

}
