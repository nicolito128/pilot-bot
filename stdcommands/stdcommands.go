package stdcommands

import (
	"github.com/nicolito128/pilot-bot/pkg/commands"
	"github.com/nicolito128/pilot-bot/stdcommands/calculator"
	"github.com/nicolito128/pilot-bot/stdcommands/info"
	"github.com/nicolito128/pilot-bot/stdcommands/ping"
)

var CommandList = []commands.PilotCommand{
	ping.Command,

	info.DevServerCommand,
	info.GithubCommand,
	info.PilotCommand,
	info.InviteCommand,

	calculator.Command,
}

var HandlerList = map[string]commands.CommandHandler{
	ping.Command.ID: ping.Command.Handler,

	info.DevServerCommand.ID: info.DevServerCommand.Handler,
	info.GithubCommand.ID:    info.GithubCommand.Handler,
	info.PilotCommand.ID:     info.PilotCommand.Handler,
	info.InviteCommand.ID:    info.InviteCommand.Handler,

	calculator.Command.ID: calculator.Command.Handler,
}
