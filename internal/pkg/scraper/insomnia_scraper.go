package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

// InsomniaScraper struct.
type InsomniaScraper struct{}

// NewInsomniaScraper returns an InsomniaScraper.
func NewInsomniaScraper() *InsomniaScraper {
	return &InsomniaScraper{}
}

// Scrape is a Insomnia Scraper.
func (dv *InsomniaScraper) Scrape(ch chan Item) {
	c := colly.NewCollector()

	c.OnHTML(".ipsDataItem", func(e *colly.HTMLElement) {
		temp := Item{}
		temp.URL = e.ChildAttr(".ipsContained a", "href")
		temp.Title = e.ChildAttr(".ipsContained a", "title")
		temp.Source = "Insomnia"

		ch <- temp
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnScraped(func(response *colly.Response) {
		return
	})
	c.Visit("https://www.insomnia.gr/forums/forum/56-%CF%80%CF%81%CE%BF%CF%83%CF%86%CE%BF%CF%81%CE%AD%CF%82/?sortby=start_date&sortdirection=desc")
}
