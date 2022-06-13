package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

var project_uuid_option_name = "project_uuid"
var VoteCommand = discordgo.ApplicationCommand{
	Name:        "vote",
	Description: "Vote your favorite project",
	Options: []*discordgo.ApplicationCommandOption{

		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        project_uuid_option_name,
			Description: "Project ID (Should be a valid UUID)",
			Required:    true,
		},
	},
}

func VoteCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	// This example stores the provided arguments in an []interface{}
	// which will be used to format the bot's response
	message := ""

	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap[project_uuid_option_name]; ok {
		// Option values must be type asserted from interface{}.
		// Discordgo provides utility functions to make this simple.
		optionValue := option.StringValue()

		if !isValidUUID(optionValue) {
			message = "You are not enter a valid Project ID"
		} else {
			message = "You are vote the project " + optionValue + ". The new total amount of votes for the project is 1."
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
