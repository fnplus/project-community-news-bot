package main

import (
	"context"

	"github.com/fnplus/community-news-bot/datastore"

	"github.com/barthr/newsapi"
)

// IWorker interface to work with multiple sources of news and put it all into
// the data store
type IWorker interface {
	GetNewsWithKeywords(keywords string) ([]newsapi.Article, error)
	PushToDB([]datastore.NewsItem) error
}

// NewWorker creates a new instance of iWorker
func NewWorker(newsAPIKey string) IWorker {
	return &Worker{newsAPIKey: newsAPIKey}
}

// Worker implements the iWorker interface
type Worker struct {
	newsAPIKey string
}

// GetNewsWithKeywords gets the news from the newsapi source, for the given
// keywords
func (w *Worker) GetNewsWithKeywords(keywords string) ([]newsapi.Article, error) {
	client := newsapi.NewClient(w.newsAPIKey)
	ar, err := client.GetEverything(context.Background(), &newsapi.EverythingParameters{
		Keywords: keywords,
		Language: "en",
	})
	if err != nil {
		return nil, err
	}
	return ar.Articles, nil
}

// PushToDB pushes the given news article(s) to the firebase firestore
func (w *Worker) PushToDB(newsItems []datastore.NewsItem) error {
	client := datastore.NewClient()
	return client.AddNewsItems(newsItems)
}