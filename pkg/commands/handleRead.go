package commands

import (
	"fmt"

	"github.com/DawidBudzynsky/reddit_reader/pkg/audio"
	redditapi "github.com/DawidBudzynsky/reddit_reader/pkg/reddit_api"
	texttospeach "github.com/DawidBudzynsky/reddit_reader/pkg/textToSpeach"
)

type ReadHandler struct {
	BaseCommand
	RedditManager redditapi.RedditManager
	TextToSoeach  texttospeach.TextToSpeach
}

// !read --subreddit twosentencehorrorstories --timeframe day

func (h *ReadHandler) Execute() {
	vs, err := h.Session.State.VoiceState(h.Message.GuildID, h.Message.Author.ID)
	if err != nil {
		fmt.Println("coulndt find voicestate ")
	}

	connection, err := h.Session.ChannelVoiceJoin(h.Message.GuildID, vs.ChannelID, false, false)
	if err != nil {
		fmt.Println("error joining a channel: \n", err)
	}

	posts, err := h.RedditManager.GetTopPosts(1, redditapi.TimeFrameAll)
	if err != nil {
		fmt.Println("error while downloading reddit stories")
	}

	// TODO: add queue here
	for _, post := range posts {
		wholePost := post.Title + post.Body

		h.TextToSoeach.TextToMp3(wholePost)

		textToSend := fmt.Sprintf("Currently reading: %s", post.URL)
		_, err := h.Session.ChannelMessageSend(h.Message.ChannelID, textToSend)
		if err != nil {
			fmt.Println("Error sending message to channel:", err)
		}
	}

	audio.PlayAudioFile(connection, "output.mp3")
}
