package commands

import (
	"fmt"

	"github.com/DawidBudzynsky/reddit_reader/pkg/audio"
	"github.com/bwmarrin/discordgo"
)

func HelloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "world!")
	}
}

func VoiceChannelJoin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "join" {
		return
	}

	vs, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("coulndt find voicestate ")
	}

	connection, err := s.ChannelVoiceJoin(m.GuildID, vs.ChannelID, false, false)
	fmt.Println(connection.Ready)
	if err != nil {
		fmt.Println("error joining a channel: \n", err)
	}
}

func PlayAudio(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "play" {
		return
	}

	vs, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("coulndt find voicestate ")
	}

	connection, err := s.ChannelVoiceJoin(m.GuildID, vs.ChannelID, false, false)
	fmt.Println(connection.Ready)
	if err != nil {
		fmt.Println("error joining a channel: \n", err)
	}

	audio.PlayAudioFile(connection, "output.mp3")
}
