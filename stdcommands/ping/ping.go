package ping

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/pilot-bot/pkg/commands"
)

var Command = commands.PilotCommand{
	ID: "ping",
	Data: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Current bot latency.",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()),
			},
		})
	},
}
