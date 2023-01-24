package commands

import "github.com/bwmarrin/discordgo"

type CommandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

type PilotCommand struct {
	ID      string
	Data    *discordgo.ApplicationCommand
	Handler CommandHandler
}
