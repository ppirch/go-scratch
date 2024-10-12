package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/ppirch/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToAPIUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
	UserID        uuid.UUID `json:"user_id"`
}

func databaseFeedToAPIFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		Url:       feed.Url,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		// LastFetchedAt: feed.LastFetchedAt,
		UserID: feed.UserID,
	}
}

func databaseFeedToAPIFeeds(feeds []database.Feed) []Feed {
	apiFeeds := make([]Feed, len(feeds))
	for i, feed := range feeds {
		apiFeeds[i] = databaseFeedToAPIFeed(feed)
	}
	return apiFeeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToAPIFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func databaseFeedFollowToAPIFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	apiFeedFollows := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		apiFeedFollows[i] = databaseFeedFollowToAPIFeedFollow(feedFollow)
	}
	return apiFeedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePostToPost(post database.Post) Post {
	var description *string
	if post.Description.Valid {
		description = &post.Description.String
	}

	return Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: description,
		PublishedAt: post.PublishedAt,
		Url:         post.Url,
		FeedID:      post.FeedID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

func databasePostsToPost(posts []database.Post) []Post {
	apiPosts := make([]Post, len(posts))
	for i, post := range posts {
		apiPosts[i] = databasePostToPost(post)
	}
	return apiPosts
}
