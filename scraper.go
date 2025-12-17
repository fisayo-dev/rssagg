package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/fisayo-dev/rssagg/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration){
	log.Printf("Scraping on %v goroutines every %s diration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	// context.Background() - This is a global context this is whe u use when u don;t have access to a scoped context.
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(),int32(concurrency))
		if err != nil{
			log.Println("error fethcing feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		// Loop throgh feed with length set to the concurrency
		for _, feed := range feeds {
			wg.Add(1)
			// spawn new go routine 
			go scrapeFeed(db,wg,feed)
		}
		// Waits for all go routines to be done.
		wg.Wait()
	} 
}

func scrapeFeed(db *database.Queries,wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(),feed.ID)
	if err != nil {
		log.Println("Error marking feeds as fetched:", err)
		return
	}
	
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetchign feeds:", err)
		return
	}

	for _,item := range rssFeed.Channel.Item {
		log.Println("Found post:", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected, %v post found",feed.Name, len(rssFeed.Channel.Item))
}