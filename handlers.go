package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func inputHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the !checkers prefix
	if !strings.HasPrefix(m.Content, "!fridge") {
		return
	}
}
