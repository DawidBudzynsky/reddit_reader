package redditapi

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type ReadablePost struct {
	Title string
	Body  string
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
}

type SubredditManager struct {
	Client    *reddit.Client
	Subreddit string
}

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

func (m *SubredditManager) changePostToReadablePost(posts []*reddit.Post) []*ReadablePost {
	readablePosts := []*ReadablePost{}
	for _, post := range posts {
		readablePosts = append(readablePosts, &ReadablePost{
			Title: post.Title,
			Body:  post.Body,
		})
	}
	return readablePosts
}
