package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var UserCommand = discordgo.ApplicationCommand{
	Name:        "user",
	Description: "Show information about user",
	Options: []*discordgo.ApplicationCommandOption{

		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "User from server",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        "amount",
			Description: "Amount",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "category",
			Description: "Category of Reward",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{Name: "Kindness", Value: "kindness"},
				{Name: "Popularity", Value: "popularity"},
				{Name: "Brilliance", Value: "brilliance"},
				{Name: "Modak Fan", Value: "modak_fan"},
				{Name: "Reference", Value: "reference"},
				{Name: "Gift", Value: "gift"},
				{Name: "Other", Value: "other"},
			},
		},
	},
}

func UserCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	Var("Options", options)
	Var("User", i.Interaction.Member)
	message := fmt.Sprintln("Called by User Id:", i.Interaction.Member.User.ID)
	message += fmt.Sprintln("Username:", i.Interaction.Member.User.String())

	for _, o := range options {
		if o.Name == "user" {
			message += fmt.Sprintln("Receiver:", o.Value.(string))
		}

		if o.Name == "category" {
			message += fmt.Sprintln("Category:", o.Value.(string))
		}

		if o.Name == "amount" {
			message += fmt.Sprintln("Amount:", int(o.IntValue()))
		}
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func Var(name string, any any) {
	json, _ := json.Marshal(any)
	log.Printf("%s: %s", strings.ToUpper(name), json)
}
