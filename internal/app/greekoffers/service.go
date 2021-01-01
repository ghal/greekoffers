package greekoffers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/ghal/greekoffers/internal/pkg/publisher"
	"github.com/ghal/greekoffers/internal/pkg/scraper"
	"github.com/ghal/greekoffers/pkg/elastic/offer"
)

// Service struct.
type Service struct {
	r *offer.Repository
}

// NewService returns a new service.
func NewService(r *offer.Repository) *Service {
	return &Service{r: r}
}

// Run runs the service.
func (s *Service) Run() {
	mapping := `{
          "mappings": {
            "_doc": {
              "properties": {
                "id":                  { "type": "keyword" },
                "url":                 { "type": "keyword" },
                "source":              { "type": "keyword" },
                "title":               { "type": "text", "analyzer": "standard" },
	            "price":               { "type": "dense_vector", "dims": "2" },
	            "discount_percentage": { "type": "integer" },
	            "@timestamp":          { "type": "date" },
              }
            }
          }
		}`
	err := s.r.CreateIndex(mapping)
	if err != nil {
		fmt.Errorf("%w", err)
	}

	ch := make(chan scraper.Item)
	go func() {
		for {
			i := Item(<-ch)
			d := &offer.Document{
				ID:                 getMD5Hash(i.URL),
				Source:             i.Source,
				Title:              strip(i.Title),
				DiscountPercentage: i.DiscountPercentage,
				URL:                i.URL,
				Price:              i.Price,
				Timestamp:          time.Now().UTC(),
			}
			if s.shouldPublish(d) {
				// save to es.
				err := s.r.Create(d)
				if err != nil {
					fmt.Println(err)
				}

				// publish document.
				Publish(d)
			}
		}
	}()
	ScrapeAllSources(ch)
}

// ScrapeAllSources runs all the scrappers on a time interval.
func ScrapeAllSources(ch chan scraper.Item) {
	scr := scraper.NewService()
	lg := scr.GetScraper(scraper.Lagonika)
	sk := scr.GetScraper(scraper.Skroutz)
	is := scr.GetScraper(scraper.Insomnia)

	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			go sk.Scrape(ch)
			go lg.Scrape(ch)
			go is.Scrape(ch)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

// shouldPublish contains the rules about publishing the item.
func (s *Service) shouldPublish(item *offer.Document) bool {
	exists, err := s.r.Exists(item)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		return false
	}

	// skroutz rules
	if item.Source == "Skroutz" {
		if item.DiscountPercentage != nil && *item.DiscountPercentage > 40 {
			return true
		}
		return false
	}

	return true
}

// Publish publishes the scraped item to the publisher channels.
func Publish(item *offer.Document) {
	scr := publisher.NewPublishers()
	tp := scr.GetPublisher(publisher.Telegram)
	tp.Publish(publisher.Item{
		Title:              item.Title,
		URL:                item.URL,
		DiscountPercentage: nil,
		Price:              item.Price,
	})
}

func strip(s string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-ZΑ-Ωα-ωΆ-Ώά-ώ0-9$€ .,\\-()%]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
