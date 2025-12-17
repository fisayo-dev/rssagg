package main

import (
	"time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ApiKey string `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		Email: dbUser.Email,
		Password: dbUser.Password,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey: dbUser.ApiKey,
	}
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url     string `json:"url"`
	UserID  uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Url: dbFeed.Url,
		Name: dbFeed.Name,
		UserID: dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func databaseFeedsToFeed(dbFeeds []database.Feed) []Feed {
	totalFeeds := []Feed{}
	for _, feed := range dbFeeds{
		totalFeeds = append(totalFeeds, databaseFeedToFeed(feed))
	}
	return totalFeeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedsFollow) FeedFollow{
	return FeedFollow{
		ID: dbFeedFollow.ID, 
		CreatedAt: dbFeedFollow.CreatedAt, 
		UpdatedAt: dbFeedFollow.UpdatedAt, 
		FeedID: dbFeedFollow.FeedID, 
		UserID: dbFeedFollow.UserID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedsFollow) []FeedFollow{
	feeds := []FeedFollow{}
	for _, feedFollow := range dbFeedFollows {
		feeds = append(feeds, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feeds
}

type Post struct{
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string `json:"title"`
	Description *string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}	

func databasePostsToPosts(dbPosts []database.Post) []Post{
	finalPosts := []Post{}
	for _, post := range dbPosts {
		finalPosts = append(finalPosts, databasePostToPost(post))
	}
	return finalPosts
} 

func databasePostToPost(dbPost database.Post) Post{
	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
} 
