package discord

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//Globals that will be used throughout this directory
var (
	Token     string
	BotPrefix string
)

func LoadConfig() error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
	Token = os.Getenv("DISCORD_BOT_TOKEN")

	BotPrefix = "$"

	return nil
}
