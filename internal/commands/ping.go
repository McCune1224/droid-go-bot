package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

// Wrap discordgo ApplicationCommand struct to be called PingCommand
type PingCommand struct {
	Details dg.ApplicationCommand
}

// PingCommand implementation of the SlashCommand Interface
// Responds to the invoker the text of "pong!"
func (pc PingCommand) Invoke(s *dg.Session, i *dg.InteractionCreate) {
	ResponseData := &dg.InteractionResponseData{
		Content: "pong!",
	}

	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: ResponseData,
	})
}

// func (pc PingCommand) Attach(contract *SlashHandlers) {
// 	contract[pc.Name] = pc.Invoke
//
// }

// Declare the Ping command for implementation
var Ping = PingCommand{
	Details: dg.ApplicationCommand{
		Name:        "ping",
		Description: "Responds with 'pong!'.",
	}}
