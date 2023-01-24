package calculator

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	calc "github.com/nicolito128/go-calculator"
	"github.com/nicolito128/pilot-bot/pkg/commands"
)

var Command = commands.PilotCommand{
	ID: "calculator",

	Data: &discordgo.ApplicationCommand{
		Name:        "calculator",
		Description: "Enter expressions and use it as a normal calculator.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "expression",
				Description: "Mathematical expression.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "decimals",
				Description: "Decimal digits to show up.",
				Required:    false,
			},
		},
	},

	Handler: CalculatorHandler,
}

func CalculatorHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var messageContent *discordgo.MessageEmbed
	var expression, decimalFormat string
	var decimals int

	if option, ok := optionMap["expression"]; ok {
		expression = option.StringValue()
	}

	if option, ok := optionMap["decimals"]; ok {
		decimals = int(option.IntValue())
	}

	switch decimals {
	case 0:
		decimalFormat = "`%.0f`"
	case 1:
		decimalFormat = "`%.1f`"
	case 2:
		decimalFormat = "`%.2f`"
	case 3:
		decimalFormat = "`%.3f`"
	case 4:
		decimalFormat = "`%.4f`"
	default:
		decimalFormat = "`%.6f`"
	}

	res, err := calc.Resolve(expression)
	if err == nil {
		messageContent = &discordgo.MessageEmbed{
			Title:       "Result",
			Description: fmt.Sprintf(decimalFormat, res),
			Color:       int(0x79daf7),
		}
	} else {
		messageContent = &discordgo.MessageEmbed{
			Title:       "Error",
			Description: fmt.Sprintf("`%s`", err.Error()),
			Color:       int(0xf78879),
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{messageContent},
		},
	})
}
