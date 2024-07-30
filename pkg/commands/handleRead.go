package commands

import (
	"fmt"
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

func (h *ReadHandler) GetNextInQueue() (string, bool) {
	h.QueueMutex.Lock()
	defer h.QueueMutex.Unlock()
	if len(h.Queue) == 0 {
		return "", false
	}
	next := h.Queue[0]
	h.Queue = h.Queue[1:]
	return next, true
}

func (h *ReadHandler) SetPlaying(isPlaying bool) {
	h.PlayingMutex.Lock()
	defer h.PlayingMutex.Unlock()
	h.IsPlaying = isPlaying
}

func (h *ReadHandler) IsCurrentlyPlaying() bool {
	h.PlayingMutex.Lock()
	defer h.PlayingMutex.Unlock()
	return h.IsPlaying
}

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

func (h *ReadHandler) Disconnect() {
	h.VoiceConnection.Disconnect()
	h.VoiceConnection.Close()
	h.VoiceConnection = nil
}

func (h *ReadHandler) Execute() {
	if h.VoiceConnection == nil {
		h.joinVoiceChannel()
	}

	posts, err := h.RedditManager.GetTopPosts(1, redditapi.TimeFrameAll)
	if err != nil {
		fmt.Println("error while downloading reddit stories")
		return
	}

	h.AddPostsToQueue(posts)
	if !h.IsCurrentlyPlaying() {
		h.playAudioQueue()
	}

	// textToSend := fmt.Sprintf("Currently reading: %s", post.URL)
	// _, err := h.Session.ChannelMessageSend(h.Message.ChannelID, textToSend)
	// if err != nil {
	// 	fmt.Println("Error sending message to channel:", err)
	// }
}

func (h *ReadHandler) AddPostsToQueue(posts []*redditapi.ReadablePost) {
	// TODO: make it run in goroutines
	for _, post := range posts {
		wholePost := post.Title + post.Body
		audioFilePath := h.generateAudioFile(wholePost)
		h.AddToQueue(audioFilePath)
	}
}

func (h *ReadHandler) AddToQueue(audioFile string) {
	h.QueueMutex.Lock()
	defer h.QueueMutex.Unlock()
	h.Queue = append(h.Queue, audioFile)
}

func (h *ReadHandler) playAudioQueue() {
	h.SetPlaying(true)
	defer h.SetPlaying(false)
	for {
		audioFile, ok := h.GetNextInQueue()
		if !ok {
			break
		}
		audio.PlayAudioFile(h.VoiceConnection, audioFile)
	}
}

func (h *ReadHandler) generateAudioFile(text string) string {
	// Use a unique filename for each audio file
	fileName := fmt.Sprintf(".output/output_%d.mp3", time.Now().UnixNano())
	h.TextToSpeech.SetDestinationPath(fileName)
	h.TextToSpeech.TextToMp3(text)
	return fileName
}
