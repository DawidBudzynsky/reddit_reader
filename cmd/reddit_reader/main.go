package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DawidBudzynsky/reddit_reader/pkg/commands"
	redditapi "github.com/DawidBudzynsky/reddit_reader/pkg/reddit_api"
	texttospeach "github.com/DawidBudzynsky/reddit_reader/pkg/textToSpeach"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}

	redditClient, err := reddit.NewClient(reddit.Credentials{}, reddit.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	session, err := discordgo.New(os.Getenv("DISCORD_BOT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	redditManager := redditapi.NewSubredditManager(redditClient, "TwoSentenceHorror")
	textToSpeach := texttospeach.NewTikTokTTS(texttospeach.VoiceEnglishUK2, ".output/")

	// Possible commands
	commandHandler := commands.NewCommandHandler("!", session)
	commandHandler.AddCommand("!hello", &commands.HelloWorldHandler{})
	commandHandler.AddCommand("!read", &commands.ReadHandler{
		RedditManager: redditManager,
		TextToSpeech:  textToSpeach,
	})

	// Discrod handlers
	session.AddHandler(commandHandler.HandleCommand)
	session.AddHandler(ready)

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGABRT, os.Interrupt)
	<-sc
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "reddit reader")
}
