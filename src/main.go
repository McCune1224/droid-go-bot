package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/McCune1224/oomdroid/internal/commands"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

// Add Slash Command Details
var (
	discordBot *discordgo.Session
)

// Load env stuff
func init() {
	var err error
	token := os.Getenv("BOT_TOKEN")
	discordBot, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	var err error

	err = discordBot.Open()
	defer discordBot.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Open Connection

	//Make Channel to listen for exit signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("%s Listening for Commands", discordBot.State.User.Username)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
