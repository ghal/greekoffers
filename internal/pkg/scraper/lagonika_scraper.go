package scraper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

// LagonikaScraper struct.
type LagonikaScraper struct{}

// NewLagonikaScraper returns an LagonikaScraper.
func NewLagonikaScraper() *LagonikaScraper {
	return &LagonikaScraper{}
}

// Scrape is a Lagonika Scraper.
func (dv *LagonikaScraper) Scrape(ch chan Item) {
	c := colly.NewCollector()

	c.OnHTML(".lagonika-offer-main", func(e *colly.HTMLElement) {
		temp := Item{}
		temp.URL = e.ChildAttr(".la-listview-title a", "href")
		temp.Title = e.ChildText(".la-listview-title")
		temp.Source = "Lagonika"
		// keep only numbers.
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		fl := re.FindAllString(e.ChildText(".la-offer-price"), -1)
		if len(fl) > 0 {
			if s, err := strconv.ParseFloat(fl[0], 64); err == nil {
				temp.Price = &s
			}
		}
		ch <- temp
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnScraped(func(response *colly.Response) {
		return
	})
	c.Visit("https://www.lagonika.gr/?dtype=prosfata")
}
