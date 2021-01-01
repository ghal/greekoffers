package publisher

import "fmt"

// StdoutPublisher struct.
type StdoutPublisher struct{}

// NewStdoutPublisher returns an StdoutPublisher.
func NewStdoutPublisher() *StdoutPublisher {
	return &StdoutPublisher{}
}

// Publish publishes the scraped item to stdout.
func (dv *StdoutPublisher) Publish(item Item) {
	fmt.Println("Title: " + item.Title)
	fmt.Println("URL: " + item.URL)
	if item.Price != nil {
		fmt.Printf("Price: %f \n", *item.Price)
	}
}
