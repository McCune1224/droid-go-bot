package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

type Feature struct {
	Type         dg.InteractionType
	Handler      func(*dg.Session, *dg.InteractionCreate)
	SlashCommand *dg.ApplicationCommand
	CustomID     string
}

type CommandManager struct {
	Features   []*Feature
	commandIDs map[string][]string
}

func (cm CommandManager) NewCommandManager() *CommandManager {
	return &CommandManager{
		Features:   []*Feature{},
		commandIDs: map[string][]string{},
	}
}

func (cm *CommandManager) AddFeature(feature *Feature) {
	cm.Features = append(cm.Features, feature)

}
func (cm *CommandManager) AddFeatures(features []*Feature) {
	cm.Features = append(cm.Features, features...)

}

func (cm *CommandManager) CreateCommands(sess *dg.Session, gID string) error {
	return nil
}
