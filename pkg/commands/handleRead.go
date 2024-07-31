package commands

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/DawidBudzynsky/reddit_reader/pkg/audio"
	redditapi "github.com/DawidBudzynsky/reddit_reader/pkg/reddit_api"
	texttospeach "github.com/DawidBudzynsky/reddit_reader/pkg/textToSpeach"
	"github.com/bwmarrin/discordgo"
)

type ReadHandler struct {
	BaseCommand
	RedditManager   redditapi.RedditManager
	TextToSpeech    texttospeach.TextToSpeach
	Queue           []string
	QueueMutex      sync.Mutex
	VoiceConnection *discordgo.VoiceConnection
	IsPlaying       bool
	PlayingMutex    sync.Mutex
}

// Execute handles the main logic for the ReadHandler, including joining the voice channel,
// fetching posts from Reddit, generating audio files, and playing them in the voice channel.
func (h *ReadHandler) Execute() {
	if h.VoiceConnection == nil {
		h.joinVoiceChannel()
	}

	subreddit, ok := h.argumentsMap["--subreddit"]
	if !ok {
		subreddit = "golang"
	}
	h.RedditManager.SetSubreddit(subreddit)

	timeFrame, ok := h.argumentsMap["--timeframe"]
	if !ok {
		timeFrame = string(redditapi.TimeFrameAll)
	}
	numberOfPosts, ok := h.argumentsMap["--number"]
	if !ok {
		numberOfPosts = "1"
	}
	num, err := strconv.Atoi(numberOfPosts)
	if err != nil {
		fmt.Errorf("error while convertin argument: --number")
	}

	posts, err := h.RedditManager.GetTopPosts(num, redditapi.TimeFrame(timeFrame))
	if err != nil {
		fmt.Println("error while downloading reddit stories")
		return
	}

	h.addPostsToQueue(posts)
	if !h.isCurrentlyPlaying() {
		h.playAudioQueue()
	}
}

// joinVoiceChannel makes the bot join the voice channel of the user who issued the command.
func (h *ReadHandler) joinVoiceChannel() {
	vs, err := h.Session.State.VoiceState(h.Message.GuildID, h.Message.Author.ID)
	if err != nil {
		fmt.Println("coulndt find voicestate ")
		return
	}
	connection, err := h.Session.ChannelVoiceJoin(h.Message.GuildID, vs.ChannelID, false, false)
	if err != nil {
		fmt.Println("error joining a channel: \n", err)
		return
	}
	h.VoiceConnection = connection
}

// addPostsToQueue converts Reddit posts to audio files and adds them to the play queue.
func (h *ReadHandler) addPostsToQueue(posts []*redditapi.ReadablePost) {
	for _, post := range posts {
		wholePost := post.Title + post.Body
		audioFilePath := h.generateAudioFile(wholePost)
		h.addToQueue(audioFilePath)

		notifyText := "Added to Queue: " + post.URL
		_, err := h.Session.ChannelMessageSend(h.ChannelID, notifyText)
		if err != nil {
			fmt.Println("Error sending message to channel:", err)
		}
	}
}

// addToQueue adds an audio file path to the play queue.
func (h *ReadHandler) addToQueue(audioFile string) {
	h.QueueMutex.Lock()
	defer h.QueueMutex.Unlock()
	h.Queue = append(h.Queue, audioFile)
}

// getNextInQueue retrieves the next audio file in the queue for playback.
func (h *ReadHandler) getNextInQueue() (string, bool) {
	h.QueueMutex.Lock()
	defer h.QueueMutex.Unlock()
	if len(h.Queue) == 0 {
		return "", false
	}
	next := h.Queue[0]
	h.Queue = h.Queue[1:]
	return next, true
}

// isCurrentlyPlaying checks if an audio file is currently being played.
func (h *ReadHandler) isCurrentlyPlaying() bool {
	h.PlayingMutex.Lock()
	defer h.PlayingMutex.Unlock()
	return h.IsPlaying
}

// setPlaying updates the playing status of the bot.
func (h *ReadHandler) setPlaying(isPlaying bool) {
	h.PlayingMutex.Lock()
	defer h.PlayingMutex.Unlock()
	h.IsPlaying = isPlaying
}

// generateAudioFile converts a given text to an audio file and returns the file path.
func (h *ReadHandler) generateAudioFile(text string) string {
	// Use a unique filename for each audio file
	fileName := fmt.Sprintf(".output/output_%d.mp3", time.Now().UnixNano())
	h.TextToSpeech.SetDestinationPath(fileName)
	h.TextToSpeech.TextToMp3(text)
	return fileName
}

// playAudioQueue continuously plays audio files from the queue until it is empty.
func (h *ReadHandler) playAudioQueue() {
	h.setPlaying(true)
	defer h.disconnect()

	// Sleep between reading files
	pauseDuration := 2 * time.Second
	for {
		audioFile, ok := h.getNextInQueue()
		if !ok {
			break
		}
		audio.PlayAudioFile(h.VoiceConnection, audioFile)
		time.Sleep(pauseDuration)

		// Remove a file after reading it
		if err := os.Remove(audioFile); err != nil {
			fmt.Printf("Error deleting file %s: %v\n", audioFile, err)
		}
	}
}

// disconnect disconnects the bot from the voice channel and updates its playing status.
func (h *ReadHandler) disconnect() {
	h.setPlaying(false)
	h.VoiceConnection.Disconnect()
	h.VoiceConnection.Close()
	h.VoiceConnection = nil
}
