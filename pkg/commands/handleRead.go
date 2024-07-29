package commands

import (
	"fmt"

	"github.com/DawidBudzynsky/reddit_reader/pkg/audio"
)

type ReadHandler struct {
	BaseCommand
}

func (h *ReadHandler) Execute() {
	vs, err := h.Session.State.VoiceState(h.Message.GuildID, h.Message.Author.ID)
	if err != nil {
		fmt.Println("coulndt find voicestate ")
	}

	connection, err := h.Session.ChannelVoiceJoin(h.Message.GuildID, vs.ChannelID, false, false)
	fmt.Println(connection.Ready)
	if err != nil {
		fmt.Println("error joining a channel: \n", err)
	}

	audio.PlayAudioFile(connection, "output.mp3")
}
