package offer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// Repository allows to index and search documents.
type Repository struct {
	es        *elasticsearch.Client
	indexName string
}

// Document wraps an offer item.
type Document struct {
	ID                 string    `json:"id"`
	Source             string    `json:"source"`
	Title              string    `json:"title"`
	URL                string    `json:"url"`
	DiscountPercentage *int      `json:"discount_percentage"`
	Price              *float64  `json:"price"`
	Timestamp          time.Time `json:"@timestamp"`
}

// NewRepository constructor.
func NewRepository(es *elasticsearch.Client, idx string) (*Repository, error) {
	return &Repository{
		es:        es,
		indexName: idx,
	}, nil
}

// StoreConfig configures the store.
type StoreConfig struct {
	Client    *elasticsearch.Client
	IndexName string
}

// CreateIndex creates a new index with mapping.
func (s *Repository) CreateIndex(mapping string) error {
	res, err := s.es.Indices.Create(
		s.indexName,
		s.es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}

// Create indexes a new document into store.
func (s *Repository) Create(item *Document) error {
	payload, err := json.Marshal(item)
	if err != nil {
		return err
	}

	ctx := context.Background()
	res, err := esapi.CreateRequest{
		Index:      s.indexName,
		DocumentID: item.ID,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, s.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s]", res.Status())
	}

	return nil
}

// Exists returns true when a document with id already exists in the store.
func (s *Repository) Exists(item *Document) (bool, error) {
	res, err := s.es.Exists(s.indexName, item.ID)
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		return false, fmt.Errorf("[%s]", res.Status())
	}
}
