package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	SetMessage(*discordgo.MessageCreate)
	SetSession(*discordgo.Session)
	SetChannelID(string)
	GetArguments() []string
	SetArgumentsMap(map[string]string)
	Execute()
}

type BaseCommand struct {
	Message           *discordgo.MessageCreate
	Session           *discordgo.Session
	ChannelID         string
	PossibleArguments []string
	argumentsMap      map[string]string
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

func (b *BaseCommand) GetArguments() []string {
	return b.PossibleArguments
}

func (b *BaseCommand) SetArgumentsMap(argumentsMap map[string]string) {
	b.argumentsMap = argumentsMap
}
