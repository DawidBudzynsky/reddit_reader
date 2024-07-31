package redditapi

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type ReadablePost struct {
	Title string
	Body  string
	URL   string
}

type TimeFrame string

const (
	TimeFrameHour  TimeFrame = "hour"
	TimeFrameDay   TimeFrame = "day"
	TimeFrameWeek  TimeFrame = "week"
	TimeFrameMonth TimeFrame = "month"
	TimeFrameYear  TimeFrame = "year"
	TimeFrameAll   TimeFrame = "all"
)

type RedditManager interface {
	GetLatestPosts(int) ([]*ReadablePost, error)
	GetTopPosts(int, TimeFrame) ([]*ReadablePost, error)
	SetSubreddit(string)
}

type SubredditManager struct {
	Client    *reddit.Client
	Subreddit string
}

func NewSubredditManager(client *reddit.Client, subreddit string) *SubredditManager {
	return &SubredditManager{
		Client:    client,
		Subreddit: subreddit,
	}
}

func (m *SubredditManager) SetSubreddit(newSubreddit string) {
	m.Subreddit = newSubreddit
}

// GetLatestPosts retrieves the latest posts from the current subreddit.
// It returns a slice of ReadablePost and any error encountered.
func (m *SubredditManager) GetLatestPosts(number int) ([]*ReadablePost, error) {
	results, _, err := m.Client.Subreddit.NewPosts(context.Background(), m.Subreddit, &reddit.ListOptions{
		Limit: number,
	})
	if err != nil {
		return nil, err
	}
	posts := m.changePostToReadablePost(results)
	return posts, nil
}

// GetTopPosts retrieves the top posts from the current subreddit based on the specified timeframe.
// It returns a slice of ReadablePost and any error encountered.
func (m *SubredditManager) GetTopPosts(number int, timeframe TimeFrame) ([]*ReadablePost, error) {
	results, _, err := m.Client.Subreddit.TopPosts(context.Background(), m.Subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: number,
		},
		Time: string(timeframe),
	})
	if err != nil {
		return nil, err
	}
	posts := m.changePostToReadablePost(results)
	return posts, nil
}

// changePostToReadablePost converts a slice of Reddit posts to a slice of ReadablePost.
// It simplifies the post structure to only include the title, body, and URL.
func (m *SubredditManager) changePostToReadablePost(posts []*reddit.Post) []*ReadablePost {
	readablePosts := []*ReadablePost{}
	for _, post := range posts {
		readablePosts = append(readablePosts, &ReadablePost{
			Title: post.Title,
			Body:  post.Body,
			URL:   post.URL,
		})
	}
	return readablePosts
}
