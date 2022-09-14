package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

func Foobar(s *dg.Session, i *dg.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionResponseData{
			Content: "Foobar!",
		},
	})
}
