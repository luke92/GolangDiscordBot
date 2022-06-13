package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Categories = []string{
	"ART",
	"CODING",
	"FASHION",
	"FOOD",
	"OTHER",
}

var Projects = []string{
	"3476e206-56b3-48e6-89bc-72e8f25906d3. New Idea. Description of idea",
}

var category_option_name = "category"

var ProjectCommand = discordgo.ApplicationCommand{
	Name:        "project",
	Description: "Show random project using or not a Valid Category (" + strings.Join(Categories, ", ") + ")",
	Options: []*discordgo.ApplicationCommandOption{

		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        category_option_name,
			Description: "Valid Category of Project",
			Required:    false,
		},
	},
}

func ProjectCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	message := ""

	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap[category_option_name]; ok {

		optionValue := option.StringValue()

		category := strings.ToUpper(optionValue)
		if !contains(Categories, category) {
			message = optionValue + " not is valid category"
		}
	}

	if message == "" {
		message = Projects[0]
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
