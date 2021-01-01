package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ghal/greekoffers/internal/app/greekoffers"
	"github.com/ghal/greekoffers/pkg/elastic"
	"github.com/ghal/greekoffers/pkg/elastic/offer"
)

func main() {
	es := setUpES()

	offerRepo, err := offer.NewRepository(es, "offers")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	s := greekoffers.NewService(offerRepo)
	s.Run()
}

func setUpES() *elasticsearch.Client {
	serviceStoreConfig := elasticsearch.Config{
		Addresses:         []string{"http://elk:9200"},
		EnableDebugLogger: true,
	}

	return elastic.NewESClient(serviceStoreConfig)
}
