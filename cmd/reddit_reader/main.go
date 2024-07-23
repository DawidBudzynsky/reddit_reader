package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	redditapi "github.com/DawidBudzynsky/reddit_reader/pkg/reddit_api"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}

	client, err := reddit.NewClient(reddit.Credentials{}, reddit.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	golangSubreddit := &redditapi.SubredditManager{
		Client:    client,
		Subreddit: "TwoSentenceHorror",
	}

	session, err := discordgo.New(os.Getenv("DISCORD_BOT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world")
		}

		if m.Content == "reddit" {
			posts, _ := golangSubreddit.GetTopPosts(1, redditapi.TimeFrameAll)

			for _, post := range posts {
				s.ChannelMessageSendTTS(m.ChannelID, post.Title)
			}
		}
	})

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGABRT, os.Interrupt)
	<-sc
}
