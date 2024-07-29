package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	SetMessage(*discordgo.MessageCreate)
	SetSession(*discordgo.Session)
	SetChannelID(string)
	Execute()
}

type BaseCommand struct {
	Message   *discordgo.MessageCreate
	Session   *discordgo.Session
	ChannelID string
}

// SetSession sets the session for the command
func (b *BaseCommand) SetSession(session *discordgo.Session) {
	b.Session = session
}

// SetMessage sets the message for the command
func (b *BaseCommand) SetMessage(message *discordgo.MessageCreate) {
	b.Message = message
}

// SetChannelID sets the channel ID for the command
func (b *BaseCommand) SetChannelID(channelID string) {
	b.ChannelID = channelID
}
