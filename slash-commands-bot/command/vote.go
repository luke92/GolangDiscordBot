package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

var project_uuid_option_name1 = "project_uuid_1"
var project_uuid_option_name2 = "project_uuid_2"
var project_uuid_option_name3 = "project_uuid_3"
var VoteCommand = discordgo.ApplicationCommand{
	Name:        "vote",
	Description: "Vote your favorite project",
	Options: []*discordgo.ApplicationCommandOption{

		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        project_uuid_option_name1,
			Description: "Project ID (Should be a valid UUID)",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        project_uuid_option_name2,
			Description: "Project ID (Should be a valid UUID)",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        project_uuid_option_name3,
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
	allOptionsIsOk := true
	optionsValue := []string{}
	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap[project_uuid_option_name1]; ok {
		// Option values must be type asserted from interface{}.
		// Discordgo provides utility functions to make this simple.
		optionValue := option.StringValue()

		if !isValidUUID(optionValue) {
			message += "The Project ID " + optionValue + "isn't a valid ID. \n"
			allOptionsIsOk = false
		} else {
			optionsValue = append(optionsValue, optionValue)
		}
	}

	if option, ok := optionMap[project_uuid_option_name2]; ok {
		// Option values must be type asserted from interface{}.
		// Discordgo provides utility functions to make this simple.
		optionValue := option.StringValue()

		if !isValidUUID(optionValue) {
			message += "The Project ID " + optionValue + "isn't a valid ID. \n"
			allOptionsIsOk = false
		} else if contains(optionsValue, optionValue) {
			message += "You cannot repeat the project ID " + optionValue + ". Enter other Project ID\n"
			allOptionsIsOk = false
		} else {
			optionsValue = append(optionsValue, optionValue)
		}
	}

	if option, ok := optionMap[project_uuid_option_name3]; ok {
		// Option values must be type asserted from interface{}.
		// Discordgo provides utility functions to make this simple.
		optionValue := option.StringValue()

		if !isValidUUID(optionValue) {
			message += "The Project ID " + optionValue + " isn't a valid ID. \n"
			allOptionsIsOk = false
		} else if contains(optionsValue, optionValue) {
			message += "You cannot repeat the project ID " + optionValue + ". Enter other Project ID\n"
			allOptionsIsOk = false
		} else {
			optionsValue = append(optionsValue, optionValue)
		}
	}

	if allOptionsIsOk {
		for _, value := range optionsValue {
			message += "You are vote the project " + value + ". The new total amount of votes for the project is 1\n"
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
