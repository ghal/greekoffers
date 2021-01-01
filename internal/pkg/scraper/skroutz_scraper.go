package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
)

// SkroutzScraper struct.
type SkroutzScraper struct{}

// NewSkroutzScraper returns an SkroutzScraper.
func NewSkroutzScraper() *SkroutzScraper {
	return &SkroutzScraper{}
}

// Scrape is a Skroutz Scraper.
func (dv *SkroutzScraper) Scrape(ch chan Item) {
	c := colly.NewCollector()
	err := c.SetProxy("socks5://tor:9050")
	if err != nil {
		fmt.Println(err)
	}

	c.OnHTML("#price_drops_index > main > div.wrapper > section > ol > li", func(e *colly.HTMLElement) {
		temp := Item{}
		temp.URL = "https://skroutz.gr" + e.ChildAttr("a", "href")
		temp.Title = e.ChildAttr("a", "title")

		re := regexp.MustCompile("[0-9]+")
		discountPercArr := re.FindAllString(e.ChildText(".pricedrop"), -1)

		if len(discountPercArr) > 0 {
			discountPrice, _ := strconv.Atoi(discountPercArr[0])
			temp.DiscountPercentage = &discountPrice
		}
		temp.Source = "Skroutz"
		if temp.Title != "" {
			ch <- temp
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnScraped(func(response *colly.Response) {
		return
	})
	c.Visit("https://www.skroutz.gr/prosfores/c/2_5_22_54_86_226_236_261_278_309_370_382_491_498_767_904_905_927_971_1052_1223_1405_1612_1780_1864_1912_2703_4142_4312/tilefwnia-photografia-video-hlektronikoi-ypologistes-pc-oikiakes-syskeues-kinhth-thlefwnia-systhmata-asfaleias-spitiou-eidi_katharismou_oikiakis_hrisis-diakosmisi-eidi-kouzinas-video-aromata-moto-lefka_idi-epipla-audio-epoxiaka-fotismos-ergaleia-makigiaz-garden-auto-siblirwnata_diatrofis-tablets-kai-accessories-peripoiisi-Gaming-Wearables-prosopiki-ugieini-iatrika-eidi-ilektrologika-aftomatismoi.html?order_by=newest")
}
