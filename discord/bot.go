package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}

func CommandValidator(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't start with $
	if !strings.HasPrefix(m.Content, BotPrefix) {
		return
	}

	// Get the arguments
	args := strings.Split(m.Content, " ")[1:]

	// Ensure non-empty command
	if len(args) == 0 {
		return
	}

	//s.ChannelMessageSend(m.ChannelID, errorMessage("Command missing", "For a list of commands type "+BotPrefix+"help"))
}

//Activate the Bot for Usage
func Run() {

	LoadConfig()
	//Pull config info

	//Create Bot
	b, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Register handler to listen for Commands
	b.AddHandler(CommandValidator)

	// Open a websocket connection to Discord and begin listening.
	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to discord", err)
		return
	}
	defer b.Close()

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
