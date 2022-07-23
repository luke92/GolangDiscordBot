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
	},
}

func UserCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	Var("Options", options)
	Var("User", i.Interaction.Member)
	message := fmt.Sprintln("Called by User Id:", i.Interaction.Member.User.ID)
	message += fmt.Sprintln("Username:", i.Interaction.Member.User.String())

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
