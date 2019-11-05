package scrappers

import (
	"fmt"
	"github.com/gocolly/colly"
)

const (
	instamartMetroMainURL = "https://instamart.ru/metro"
)

type Product struct {

}

func ScrapInstamartLinksMetroPage() ([]string, error) {
	fmt.Println("scrapping instamart metro links")

	c := colly.NewCollector()
	links := make([]string, 0)
	// Find and visit all links
	c.OnHTML(".show-all", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		links = append(links, href)
	})

	// here we are getting the links list
	err := c.Visit(instamartMetroMainURL)
	if err != nil {
		fmt.Printf("error visiting, err: %s\n", err.Error())
		return []string{}, err
	}

	return links, nil
}

