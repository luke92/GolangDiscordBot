package command

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Var(name string, any any) {
	json, _ := json.Marshal(any)
	log.Printf("%s: %s", strings.ToUpper(name), json)
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
	}

	return roles
}

func getRoleById(s *discordgo.Session, roleID string) (*discordgo.Role, error) {
	role, err := s.State.Role(GuildID, roleID)
	if err != nil {
		Var("Error getting role", err)
		return nil, err
	}

	return role, nil
}

func addRole(s *discordgo.Session, roleID string, userID string) string {
	err := s.GuildMemberRoleAdd(GuildID, userID, roleID)
	if err != nil {
		return fmt.Sprintln("Error adding role: ", err.Error())
	}

	roleName := roleID

	role, err := getRoleById(s, roleID)
	if err == nil {
		roleName = role.Name
	}

	return fmt.Sprintln("Role ", roleName, " added succesfully")
}

func removeRole(s *discordgo.Session, roleID string, userID string) string {
	err := s.GuildMemberRoleRemove(GuildID, userID, roleID)
	if err != nil {
		return fmt.Sprintln("Error removing role: ", err.Error())
	}

	roleName := roleID
	role, err := s.State.Role(GuildID, roleID)
	if err != nil {
		Var("Error getting role", err)
	} else {
		roleName = role.Name
	}
	return fmt.Sprintln("Role ", roleName, " removed succesfully")
}

func sendDMMessage(s *discordgo.Session, i *discordgo.InteractionCreate, userID string, content string) {
	channel, err := s.UserChannelCreate(userID)
	if err != nil {
		Var("Error Creating Channel", err)
	}

	message, err := s.ChannelMessageSend(channel.ID, content)
	if err != nil {
		Var("error sending DM message", err)
		printMessage(s, i, "Failed to send you a DM. Did you disable DM in your privacy settings?", true)
	}

	Var("DM Message", message)
}

func getMembers(s *discordgo.Session) ([]*discordgo.Member, error) {
	maxLimitMembers := 1000
	lastMemberID := ""
	membersReturn := []*discordgo.Member{}

	for {
		members, err := s.GuildMembers(GuildID, lastMemberID, maxLimitMembers)
		if err != nil {
			Var("Error get Members", err)
			return membersReturn, err
		}

		membersReturn = append(membersReturn, members...)

		if len(members) < maxLimitMembers {
			break
		}

		lastMemberID = members[len(members)-1].User.ID
	}

	return membersReturn, nil
}
