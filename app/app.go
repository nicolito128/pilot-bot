package app

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/pilot-bot/stdcommands"
)

const MAX_CONCURRENT_HANDLERS = 10

var options struct {
	BotToken       string
	GuildId        string
	RemoveCommands bool
}

var Session *discordgo.Session

func Start() {
	flag.StringVar(&options.BotToken, "token", "", "Bot access token")
	flag.StringVar(&options.GuildId, "guild", "", "Test guild ID. If not passed - bot registers commands globally")
	flag.BoolVar(&options.RemoveCommands, "rcm", true, "Remove all commands after shutdowning or not")
	flag.Parse()

	if len(flag.Args()) < 1 {
		token, exists := os.LookupEnv("BOT_TOKEN")
		if exists {
			options.BotToken = token
		}

		guildId, exists := os.LookupEnv("TEST_GUILD_ID")
		if exists {
			options.GuildId = guildId
		}
	}

	s, err := discordgo.New("Bot " + options.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	Session = s

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	log.Println("Adding commands...")
	register := make([]*discordgo.ApplicationCommand, len(stdcommands.CommandList))
	for i, v := range stdcommands.CommandList {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, options.GuildId, v.Data)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Data.Name, err)
		}

		register[i] = cmd
	}

	Session.AddHandler(interactionCreateHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if options.RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range register {
			err := s.ApplicationCommandDelete(s.State.User.ID, options.GuildId, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}

func interactionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handler, ok := stdcommands.HandlerList[i.ApplicationCommandData().Name]
	guard := make(chan struct{}, MAX_CONCURRENT_HANDLERS)

	if ok {
		go func() {
			handler(s, i)
			<-guard
		}()
	}
}
