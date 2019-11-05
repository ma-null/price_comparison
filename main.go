package main

import (
	"./scrappers"
	"fmt"
)

func main() {
	links, err := scrappers.ScrapInstamartLinksMetroPage()
	if err != nil {
		return
	}
	for _, link := range links {
		products, err := obrabotchik()
	}
}

