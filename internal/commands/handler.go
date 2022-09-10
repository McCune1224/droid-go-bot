package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

type SlashCommand interface {
	Invoke(s *dg.Session, i *dg.InteractionCreate)
}

// Map to pair SlashCommand interface with a string key with the text to invoke the function
type SlashHandlers map[string]func(SlashCommand)
