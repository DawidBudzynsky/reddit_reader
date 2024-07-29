package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix   string
	session  *discordgo.Session
	message  *discordgo.MessageCreate
	commands map[string]Command
}

func NewCommandHandler(prefix string, session *discordgo.Session) *CommandHandler {
	return &CommandHandler{
		prefix:   prefix,
		session:  session,
		commands: make(map[string]Command),
	}
}

func (c *CommandHandler) AddCommand(keyword string, command Command) {
	c.commands[keyword] = command
}

func (c *CommandHandler) HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, c.prefix) {
		return
	}

	args := strings.Fields(m.Content)
	command, exists := c.commands[args[0]]
	if !exists {
		s.ChannelMessageSend(m.ChannelID, "no such command")
		return
	}

	command.SetSession(s)
	command.SetMessage(m)
	command.SetChannelID(m.ChannelID)
	go command.Execute()
}
