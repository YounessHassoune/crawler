package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
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

	writeUrlsToJson("urls.json", urls)

	fmt.Println("Number of URLs collected:", len(urls))

}

func getAllSiteMapUrls(c *colly.Collector, sitemapLink string) ([]string, error) {
	urls := []string{}
	visited := make(map[string]bool) // Map to store visited URLs

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

		if !strings.HasPrefix(link, "/tools/sitemap?") && !visited[link] {
			urls = append(urls, link)
			visited[link] = true
		}
	})

	err := c.Visit(sitemapLink)

	if err != nil {
		return nil, err
	}

	return urls, nil

}

func writeUrlsToJson(filename string, strings []string) error {
	jsonData, err := json.Marshal(strings)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
