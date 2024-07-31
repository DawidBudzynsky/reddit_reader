package commands

import (
	"fmt"
	"log"
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

// AddCommand adds a new command to the CommandHandler with a specified keyword.
// The keyword is used to match incoming messages to the correct command.
func (c *CommandHandler) AddCommand(keyword string, command Command) {
	c.commands[keyword] = command
}

// HandleCommand processes incoming messages to check if they match any registered commands.
// If the message starts with the specified prefix and matches a command keyword,
// it parses arguments and executes the command.
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
	commandArguments, err := c.parseArguments(command, args[1:])

	fmt.Println(commandArguments)

	if err != nil {
		log.Fatal("argument parser error:", err)
	}
	command.SetArgumentsMap(commandArguments)
	command.SetSession(s)
	command.SetMessage(m)
	command.SetChannelID(m.ChannelID)
	go command.Execute()
}

// parseArguments extracts command-line arguments from the message and validates them.
// It populates a map of argument names to their values based on the command's expected arguments.
// Returns the map of arguments and any error encountered during parsing.
// NOTE: its a temporary solution for arguments handling
func (c *CommandHandler) parseArguments(command Command, args []string) (map[string]string, error) {
	argumentsMap := make(map[string]string)
	if len(command.GetArguments()) == 0 {
		return nil, nil
	}
	commandArgumentsSet := make(map[string]struct{})
	for _, arg := range command.GetArguments() {
		commandArgumentsSet[arg] = struct{}{}
	}
	for _, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			return nil, fmt.Errorf("Argument %s, doesnt have '--' prefix", arg)
		}
		parts := strings.SplitN(arg[2:], "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("Argument %s, doenst have a value", arg)
		}
		key := "--" + parts[0]
		value := parts[1]
		// Check if the key is a valid command argument
		if _, exists := commandArgumentsSet[key]; exists {
			argumentsMap[key] = value
		}
	}
	return argumentsMap, nil
}
