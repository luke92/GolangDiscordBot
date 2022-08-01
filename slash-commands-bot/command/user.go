package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID      string
	dmPermission = true //If is false Only appear the command for a Admin User
	//defaultMemberPermissions int64 = discordgo.PermissionManageServer
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
	DefaultPermission: &dmPermission,
}

func UserCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	message := ""

	Var("Options", options)
	member := i.Interaction.Member
	Var("User", member)
	roles := getRoles(s)

	receiverUserID := ""
	senderUserID := member.User.ID
	for _, o := range options {
		if o.Name == "user" {
			Var("Receiver User Option", o)
			receiverUserID = o.Value.(string)
		}
	}

	roleIDs := getRoleIDs(roles)
	message += getMessageDataFromUser(s, "Sender User", senderUserID, roleIDs)
	message += getMessageDataFromUser(s, "Receiver User", receiverUserID, roleIDs)

	printMessage(s, i, message, true)

}

func getMessageDataFromUser(s *discordgo.Session, fromUser string, userID string, roleIDs []string) string {
	message := ""
	user, err := s.User(userID)
	if err != nil {
		Var(fromUser+" Error", err)
	} else {
		Var(fromUser, user)
		message += fmt.Sprintln("----------------")
		message += fmt.Sprintln(fromUser)
		message += fmt.Sprintln("----------------")
		message += fmt.Sprintln("User Id:", userID)
		message += fmt.Sprintln("Username:", user.String())
		message += fmt.Sprintln("IsBot:", user.Bot)
	}
	guildMember, err := s.GuildMember(GuildID, userID)
	if err != nil {
		Var(fromUser+" Guild Member Error", err)
	} else {
		Var(fromUser+" Guild Member", &guildMember)
		intersectionRoles := intersectionArraysString(roleIDs, guildMember.Roles)
		message += fmt.Sprintln("Count of roles in the server:", len(intersectionRoles))
	}
	return message
}

func getRoleIDs(roles []*discordgo.Role) []string {
	roleIDs := []string{}
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}
	return roleIDs
}

func getRoles(s *discordgo.Session) []*discordgo.Role {
	roles, err := s.GuildRoles(GuildID)
	if err != nil {
		Var("Roles Error", err)
	} else {
		Var("Roles", roles)
		return roles
	}
	return []*discordgo.Role{}
}

func Var(name string, any any) {
	json, _ := json.Marshal(any)
	log.Printf("%s: %s", strings.ToUpper(name), json)
}

func printMessage(s *discordgo.Session, i *discordgo.InteractionCreate, message string, isPrivate bool) {
	var flag uint64
	if isPrivate {
		flag = uint64(discordgo.MessageFlagsEphemeral)
	}
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   flag,
			Content: message,
		},
	})
}

func intersectionArraysString(a, b []string) []string {
	var s []string

	for _, m := range a {

		ok := false
		for _, n := range b {
			if m == n {
				ok = true
				break
			}
		}
		if ok {
			s = append(s, m)
		}

	}

	return s
}
