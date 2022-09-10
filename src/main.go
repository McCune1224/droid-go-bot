package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/McCune1224/oomdroid/internal/commands"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

// Bot/Session settings
var ()

var discord *discordgo.Session
var CommandHandler commands.SlashHandlers

func init() {
	var err error
	token := os.Getenv("BOT_TOKEN")
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	registeredCommands := []discordgo.ApplicationCommand{}

	registeredCommands = append(registeredCommands, commands.Ping.Details)

    discord.ApplicationCommandCreate(discord.State.User.ID, "", registeredCommands[0])
	//Open Connection
	err := discord.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer discord.Close()

	//Make Channel to listen for exit signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("%s Listening for Commands", discord.State.User.Username)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
