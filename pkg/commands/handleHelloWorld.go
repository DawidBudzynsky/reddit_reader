package commands

import (
	"fmt"
)

type HelloWorldHandler struct {
	BaseCommand
}

func (h *HelloWorldHandler) Execute() {
	_, err := h.Session.ChannelMessageSend(h.ChannelID, "world!")
	if err != nil {
		fmt.Println("Error sending message to channel:", err)
	}
}
