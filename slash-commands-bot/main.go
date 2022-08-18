package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/luke92/GolangDiscordBot/slash-commands-bot/command"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var (
	commands = []*discordgo.ApplicationCommand{
		&command.VoteCommand,
		&command.ProjectCommand,
		&command.GrantRoleCommand,
		&command.UserCommand,
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		command.VoteCommand.Name:      command.VoteCommandHandler,
		command.ProjectCommand.Name:   command.ProjectCommandHandler,
		command.GrantRoleCommand.Name: command.GrantRoleCommandHandler,
		command.UserCommand.Name:      command.UserCommandHandler,
	}
)

var s *discordgo.Session

func init() {
	receiveArguments()

	intializeDiscordBotSession()

	addCommandHandlers()
}

func receiveArguments() {
	flag.Parse()
}

func intializeDiscordBotSession() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func addCommandHandlers() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func login() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}

func openSession() {
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
}

func runBot() {

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}

func main() {
	command.GuildID = *GuildID
	login()

	openSession()

	runBot()

	log.Println("Gracefully shutting down.")
}
