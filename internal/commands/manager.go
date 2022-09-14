package commands

import (
	"log"

	dg "github.com/bwmarrin/discordgo"
)

// Wrapper Structure to couple A ApplicationCommand with a dedicated Handler
type SlashCommand struct {
	//to avoid naming stuttering (SlashCommand.Command....) discordgo.ApplicationCommand is aliased to "Feature"
	Feature dg.ApplicationCommand

	//Slash Command Handler aka how the command should operate
	Handler func(s *dg.Session, i *dg.InteractionCreate)
}

// Slash Command Manager, binds all commands to a map, the command name now is used to reference a SlashCommand
type SCM struct {
	SlashCommands map[string]SlashCommand
}

// Create New Instance of a Slash Command manager
func NewSCM() *SCM {
	scm := &SCM{
		SlashCommands: make(map[string]SlashCommand),
	}
	return scm
}

func (scm SCM) AddCommand(sc SlashCommand) {
	scm.SlashCommands[sc.Feature.Name] = sc
}

// From the map of SlashCommands stored in the SCM, make them available for the current session.
// Requires Discord Websocket Connection to be open first before calling this function

// Returns a tally of successful applications created for the session
func (scm SCM) RegisterCommands(session *dg.Session) int {
	tally := 0
	for _, slashcmd := range scm.SlashCommands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, "", &slashcmd.Feature)
		if err != nil {
			log.Printf("Failed to add command %s\nError: %s", slashcmd.Feature.Name, err.Error())
			continue
		}
		tally++
	}
	return tally
}
