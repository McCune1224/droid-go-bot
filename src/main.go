package main

import (
	"log"
	"os"
	"os/signal"

	// "github.com/McCune1224/oomdroid/internal/commands"
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

var slashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Returns 'Pong!'",
	},
	{
		Name:        "foobar",
		Description: "Returns 'Foobar'",
	},
}

var slashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
	"foobar": commands.Foobar}

func init() {
	discordBot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := slashCommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(discordBot, i)
		}
	})
}

func main() {
	var err error

	err = discordBot.Open()
	defer discordBot.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(slashCommands))
	for i, v := range slashCommands {
		cmd, err := discordBot.ApplicationCommandCreate(discordBot.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	log.Printf("Added %d command(s)", len(registeredCommands))
	//Open Connection

	//Make Channel to listen for exit signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("%s Listening for Commands", discordBot.State.User.Username)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
