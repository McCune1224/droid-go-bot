package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

var Ping = SlashCommand{
	Feature: dg.ApplicationCommand{
		Name:        "ping",
		Description: "Responds with 'Pong!'",
	},
	Handler: func(s *dg.Session, i *dg.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
			Type: dg.InteractionResponseChannelMessageWithSource,
			Data: &dg.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
}
