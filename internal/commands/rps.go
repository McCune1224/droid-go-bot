package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

/*
Return values:
Player Win = 1
Player Lose = 0
Draw = -1
Invalid Option = -2
*/
func isPlayerWinner(botChoice, playerChoice string) int {
	botChoice = strings.ToLower(botChoice)
	playerChoice = strings.ToLower(playerChoice)
	switch playerChoice {
	case "rock":
		if botChoice == "rock" {
			return -1
		} else if botChoice == "paper" {
			return 0
		} else if botChoice == "scissors" {
			return 1
		}
		return -2

	case "paper":
		if botChoice == "rock" {
			return 1
		} else if botChoice == "paper" {
			return -1
		} else if botChoice == "scissors" {
			return 0
		}
		return -2
	case "scissors":
		if botChoice == "rock" {
			return 0
		} else if botChoice == "paper" {
			return 1
		} else if botChoice == "scissors" {
			return -1
		}
		return -2
	default:
		return -2

	}

}

var Rps = SlashCommand{
	Feature: dg.ApplicationCommand{
		Name:        "rps",
		Description: "Play Rock Paper Scissors!",
		Options: []*dg.ApplicationCommandOption{
			{
				Type:        dg.ApplicationCommandOptionString,
				Name:        "selection",
				Description: "Input either Rock, Paper, or Scissors",
				Required:    true,
			},
		},
	},

	Handler: func(s *dg.Session, i *dg.InteractionCreate) {
		// botName := s.State.User.Username
		// playerName := i.User.Username
		playeSelected := i.ApplicationCommandData().Options[0].StringValue()

		rand.Seed(time.Now().UnixNano())
		Options := []string{"rock", "paper", "scissors"}
		botSelected := Options[rand.Intn(3)]

		var winner string
		botResponse := fmt.Sprintf("You Selected *%s*, I selected *%s*\n", playeSelected, botSelected)
		switch isPlayerWinner(botSelected, playeSelected) {
		case 1:
			winner = "**You Win!**"
		case 0:
			winner = "**You Lose!**"
		case -1:
			winner = "**Draw!**"
		case -2:
			winner = fmt.Sprintf("**No One Wins. Invalid option '%s'.**", playeSelected)
			botResponse = ""
		}

		botResponse += fmt.Sprintf(winner)
		s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
			Type: dg.InteractionResponseChannelMessageWithSource,
			Data: &dg.InteractionResponseData{
				Content: botResponse,
			},
		})

	},
}
