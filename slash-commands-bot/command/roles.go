package command

import "github.com/bwmarrin/discordgo"

var roles = []string{
	"3476e206-56b3-48e6-89bc-72e8f25906d3. New Idea. 2",
}

var GrantRoleCommand = discordgo.ApplicationCommand{
	Name:        "grant_role",
	Description: "Assign Role to user",
}

func GrantRoleCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	printMessage(s, i, roles[0], false)
}
