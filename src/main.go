package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/McCune1224/oomdroid/internal/commands"
	dg "github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

// Create Bot Session
var (
	discordBot *dg.Session
)

func init() {
	var err error
	token := os.Getenv("BOT_TOKEN")
	discordBot, err = dg.New("Bot " + token)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Create Command Handler and Register Slash Commands
var oomSCM *commands.SCM

func init() {
	oomSCM = commands.NewSCM()
	oomSCM.AddCommand(commands.Ping)
	oomSCM.AddCommand(commands.Rps)
	//Create handler for interactionCreation (required for Responding to slash commands)
	discordBot.AddHandler(func(s *dg.Session, i *dg.InteractionCreate) {
		log.Print(i.ApplicationCommandData().Name)
		if cmd, ok := oomSCM.SlashCommands[i.ApplicationCommandData().Name]; ok {
			cmd.Handler(s, i)
		}
	})
}

// Create Websocket Connection, Register Commands, and run until signal close
func main() {
	var err error

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err.Error())
	}
	regCommandTally := oomSCM.RegisterCommands(discordBot)
	log.Printf("Successfully registered %d/%d commands", len(oomSCM.SlashCommands), regCommandTally)
	defer discordBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("%s Listening for Commands", discordBot.State.User.Username)
	<-stop
}
