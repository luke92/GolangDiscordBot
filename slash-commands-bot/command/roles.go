package command

import "github.com/bwmarrin/discordgo"

var (
	RoleVerifiedTeen = string("1009478529899049062")
	userOption       = "user"
	roleOption       = "role"
)

var GrantRoleCommand = discordgo.ApplicationCommand{
	Name:        "grant_role",
	Description: "Assign Role to user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        userOption,
			Description: "User from server",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionRole,
			Name:        roleOption,
			Description: "Role for assign",
			Required:    true,
		},
	},
}

func GrantRoleCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	userValue := ""
	roleValue := ""

	for _, o := range options {
		if o.Name == userOption {
			Var("Receiver User Option", o)
			userValue = o.Value.(string)
		}

		if o.Name == roleOption {
			roleValue = o.Value.(string)
		}
	}

	message := addRole(s, roleValue, userValue)

	printMessage(s, i, message, false)
}

var RemoveRoleCommand = discordgo.ApplicationCommand{
	Name:        "remove_role",
	Description: "Remove Role to user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        userOption,
			Description: "User from server",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionRole,
			Name:        roleOption,
			Description: "Role for remove",
			Required:    true,
		},
	},
}

func RemoveRoleCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	userValue := ""
	roleValue := ""

	for _, o := range options {
		if o.Name == userOption {
			Var("Receiver User Option", o)
			userValue = o.Value.(string)
		}

		if o.Name == roleOption {
			roleValue = o.Value.(string)
		}
	}

	message := removeRole(s, roleValue, userValue)

	printMessage(s, i, message, false)
}
