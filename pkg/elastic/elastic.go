package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewESClient returns a handle to a Database.
func NewESClient(c elasticsearch.Config) (db *elasticsearch.Client) {
	es, err := elasticsearch.NewClient(c)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	return es
}
