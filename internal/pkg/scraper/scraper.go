package scraper

const (
	// Lagonika const
	Lagonika = "lagonika"
	// Skroutz const
	Skroutz = "skroutz"
	// Insomnia const
	Insomnia = "insomnia"
)

// Item is a scraper Item.
type Item struct {
	Title              string
	URL                string
	Source             string
	DiscountPercentage *int
	Price              *float64
}

// Scraper is a scraper interface.
type Scraper interface {
	Scrape(ch chan Item)
}

// Scrapers contains the available scrapers.
type Scrapers struct {
	LagonikaScraper *LagonikaScraper
	SkroutzScraper  *SkroutzScraper
	InsomniaScraper *InsomniaScraper
}

// NewService creates a new Scrapper service.
func NewService() Scrapers {
	return Scrapers{
		NewLagonikaScraper(),
		NewSkroutzScraper(),
		NewInsomniaScraper(),
	}
}

// GetScraper returns a new Scraper based on a provided name.
func (v *Scrapers) GetScraper(scraperName string) Scraper {
	switch scraperName {
	case Lagonika:
		return v.LagonikaScraper
	case Skroutz:
		return v.SkroutzScraper
	case Insomnia:
		return v.InsomniaScraper
	default:
		return nil
	}
}
