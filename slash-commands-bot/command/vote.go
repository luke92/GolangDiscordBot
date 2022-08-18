package command

import (
	"github.com/bwmarrin/discordgo"
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
	value := ""
	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap[project_uuid_option_name1]; ok {
		value = option.StringValue()
		messageValidation, isValid := validateOption(value, optionsValue)
		message += messageValidation
		allOptionsIsOk = allOptionsIsOk && isValid
		if isValid {
			optionsValue = append(optionsValue, value)
		}
	}

	if option, ok := optionMap[project_uuid_option_name2]; ok {
		value = option.StringValue()
		messageValidation, isValid := validateOption(value, optionsValue)
		message += messageValidation
		allOptionsIsOk = allOptionsIsOk && isValid
		if isValid {
			optionsValue = append(optionsValue, value)
		}
	}

	if option, ok := optionMap[project_uuid_option_name3]; ok {
		value = option.StringValue()
		messageValidation, isValid := validateOption(value, optionsValue)
		message += messageValidation
		allOptionsIsOk = allOptionsIsOk && isValid
		if isValid {
			optionsValue = append(optionsValue, value)
		}
	}

	if allOptionsIsOk {
		for _, value := range optionsValue {
			message += "You are vote the project " + value + ". The new total amount of votes for the project is 1\n"
		}
	}

	printMessage(s, i, message, false)
}

func validateOption(value string, optionsValue []string) (string, bool) {
	if !isValidUUID(value) {
		message := "The Project ID " + value + " isn't a valid ID. \n"
		return message, false
	} else if contains(optionsValue, value) {
		message := "You cannot repeat the project ID " + value + ". Enter other Project ID\n"
		return message, false
	} else {
		return "", true
	}
}
