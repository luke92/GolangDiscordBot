package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var CountRolesCommand = discordgo.ApplicationCommand{
	Name:        "count-roles",
	Description: "Get Information About Roles",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionRole,
			Name:        roleOption,
			Description: "Role for search",
			Required:    false,
		},
	},
}

func CountRolesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	message := ""
	roleValue := ""

	for _, o := range options {
		if o.Name == roleOption {
			roleValue = o.Value.(string)
		}
	}

	roles := []*discordgo.Role{}

	if roleValue != "" {
		role, err := getRoleById(s, roleValue)
		if err != nil {
			printMessage(s, i, err.Error(), false)
			return
		}

		roles = append(roles, role)
	} else {
		roles = getRoles(s)
		if roles == nil {
			printMessage(s, i, "Cannot get roles", false)
			return
		}
	}

	members, err := getMembers(s)
	if err != nil {
		printMessage(s, i, err.Error(), false)
		return
	}

	Var("Count Members", len(members))

	for _, role := range roles {
		message += fmt.Sprintf("Role Name: %s\n", role.Name)
		count := 0
		for _, member := range members {
			if contains(member.Roles, role.ID) {
				count++
			}
		}

		message += fmt.Sprintf("Count: %d\n\n", count)
	}

	printMessage(s, i, message, false)
}
