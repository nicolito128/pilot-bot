package info

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/pilot-bot/pkg/commands"
)

const (
	devserver = "https://discord.gg/xMybfjpQAy"
	github    = "https://github.com/nicolito128/pilot-bot"
	invite    = "https://discord.com/oauth2/authorize?client_id=1067284684930809917&scope=bot+applications.commands&permissions=1644369407095applications.commands"
)

var DevServerCommand = commands.PilotCommand{
	ID: "devserver",
	Data: &discordgo.ApplicationCommand{
		Name:        "devserver",
		Description: "Get the link to the development server.",
	},
	Handler: DevServerHandler,
}

func DevServerHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**Development discord server**: %s", devserver),
		},
	})
}

var GithubCommand = commands.PilotCommand{
	ID: "github",
	Data: &discordgo.ApplicationCommand{
		Name:        "github",
		Description: "Get the link to the github repository.",
	},
	Handler: GithubHandler,
}

func GithubHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**Github repository**: %s", github),
		},
	})
}

var InviteCommand = commands.PilotCommand{
	ID: "invite",
	Data: &discordgo.ApplicationCommand{
		Name:        "invite",
		Description: "Get a link to invite me to your server!",
	},
	Handler: InviteHandler,
}

func InviteHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**Invitation link**: %s", invite),
		},
	})
}

var PilotCommand = commands.PilotCommand{
	ID: "pilot",
	Data: &discordgo.ApplicationCommand{
		Name:        "pilot",
		Description: "Shows Pilot general information",
	},
	Handler: PilotBotDataHandler,
}

func PilotBotDataHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := generateServerEmbed(s, i.Message)
	if embed == nil {
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Embeds:  []*discordgo.MessageEmbed{embed},
		},
	})
}

func generateServerEmbed(s *discordgo.Session, m *discordgo.Message) *discordgo.MessageEmbed {
	var channels, guilds, users int
	guilds = len(s.State.Guilds)

	for _, guild := range s.State.Guilds {
		channels += len(guild.Channels)
		users += len(guild.Members)
	}

	return &discordgo.MessageEmbed{
		Title:       "Pilot",
		Description: "Information about me and how many things I'm doing.",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Guilds", Value: fmt.Sprintf("%d", guilds), Inline: true},
			{Name: "Channels", Value: fmt.Sprintf("%d", channels), Inline: true},
			{Name: "Members", Value: fmt.Sprintf("%d", users), Inline: true},
			{Name: "Ping", Value: fmt.Sprintf("%dms", s.HeartbeatLatency().Milliseconds()), Inline: true},
			{Name: "OS", Value: runtime.GOOS, Inline: true},
			{Name: "Go Version", Value: runtime.Version(), Inline: true},
			{Name: "Goroutines", Value: fmt.Sprintf("%d", runtime.NumGoroutine()), Inline: true},
			{Name: "CPU Available", Value: fmt.Sprintf("%d", runtime.NumCPU()), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Repository: https://github.com/nicolito128/pilot-bot",
		},
	}
}
