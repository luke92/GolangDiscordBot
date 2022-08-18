package command

import "github.com/bwmarrin/discordgo"

var RoleVerifiedTeen = string("1009478529899049062")

var GrantRoleCommand = discordgo.ApplicationCommand{
	Name:        "grant_role",
	Description: "Assign Role to user",
}

func GrantRoleCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	printMessage(s, i, RoleVerifiedTeen, false)
}
