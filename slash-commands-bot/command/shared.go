package command

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func printMessageEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, message string, isPrivate bool) {
	var flag uint64
	if isPrivate {
		flag = uint64(discordgo.MessageFlagsEphemeral)
	}

	pngFromFile, _ := getImageFile("image.png")
	fileImageURL := getFileImageURL("https://www.kindpng.com/picc/m/199-1998580_5-dollars-hd-png-download.png")

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Files: []*discordgo.File{
				{
					Name:        "image1.png",
					ContentType: "image/png",
					Reader:      pngFromFile,
				},
				fileImageURL,
			},
			Embeds: []*discordgo.MessageEmbed{
				{
					Color: 39423,
					Footer: &discordgo.MessageEmbedFooter{
						IconURL: "https://i.imgur.com/AfFp7pu.png",
						Text:    "Daily",
					},
					Author: &discordgo.MessageEmbedAuthor{
						Name:    i.Member.User.String(),
						IconURL: "https://i.imgur.com/AfFp7pu.png",
					},
					URL:         "https://w7.pngwing.com/pngs/737/804/png-transparent-ultron-logo-marvel-comics-chitauri-hydra-ultron-face-fictional-characters-head.png",
					Type:        discordgo.EmbedTypeImage,
					Title:       "Ultron",
					Description: fmt.Sprintf("**%s**", message),
					Image: &discordgo.MessageEmbedImage{
						URL:      "https://w7.pngwing.com/pngs/737/804/png-transparent-ultron-logo-marvel-comics-chitauri-hydra-ultron-face-fictional-characters-head.png",
						Width:    100,
						Height:   100,
						ProxyURL: "https://w7.pngwing.com/pngs/737/804/png-transparent-ultron-logo-marvel-comics-chitauri-hydra-ultron-face-fictional-characters-head.png",
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "name 1",
							Value:  "Value 1",
							Inline: false,
						},
						{
							Name:   "name 2",
							Value:  "value 2",
							Inline: true,
						},
						{
							Name:   "name 3",
							Value:  "Value 3",
							Inline: true,
						},
					},
				},
			},

			Flags:   flag,
			Content: message,
		},
	})
}

func getImageFile(filePath string) (io.Reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return f, nil
}

func getFileImageURL(url string) *discordgo.File {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &discordgo.File{
		Name:        "image2.png",
		ContentType: "image/png",
		Reader:      res.Body,
	}
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
