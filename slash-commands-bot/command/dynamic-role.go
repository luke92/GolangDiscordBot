package command

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var DynamicRoleUserCommand = discordgo.ApplicationCommand{
	Name:        "dynamic-role",
	Description: "Add Dynamic role for user for each day",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "User from server",
			Required:    true,
		},
	},
	DefaultPermission: &dmPermission,
}

func DynamicRoleUserCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	message := ""

	Var("Options", options)
	member := i.Interaction.Member
	Var("User", member)

	receiverUserID := ""

	for _, o := range options {
		if o.Name == "user" {
			Var("Receiver User Option", o)
			receiverUserID = o.Value.(string)
		}
	}

	role, err := getDynamicRole(s)

	if err != nil {
		message += "Error adding role"
		printMessage(s, i, message, true)
	}

	message = addRole(s, role.ID, receiverUserID)

	printMessage(s, i, message, false)

}

func getDynamicRole(s *discordgo.Session) (*discordgo.Role, error) {
	var err error
	roleName := getRoleNameOfTheDay()
	roles := getRoles(s)

	role := getRole(roleName, roles)

	if role == nil {
		role, err = s.GuildRoleCreate(GuildID)

		if err != nil {
			Var("Add Role Error", err)
			return nil, err
		}

		role.Name = roleName
		role.Color = 15158332

		role, err = s.GuildRoleEdit(GuildID, role.ID, role.Name, role.Color, role.Hoist, role.Permissions, role.Mentionable)

		if err != nil {
			Var("Edit Role Error", err)
			return nil, err
		}

		Var("New Role", role)
	}

	return role, nil
}

func getRoleNameOfTheDay() string {
	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day()

	yearStr := fmt.Sprint(year)
	yearStr = yearStr[len(yearStr)-2:]

	name := fmt.Sprintf("Auto-%s%02d%02d", yearStr, month, day)

	return name
}

func getRole(roleName string, roles []*discordgo.Role) *discordgo.Role {
	for _, role := range roles {
		if role.Name == roleName {
			return role
		}
	}
	return nil
}
