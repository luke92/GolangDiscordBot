package command

import (
	"fmt"
	"math/big"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID      string
	dmPermission = true //If is false Only appear the command for a Admin User
	//defaultMemberPermissions int64 = discordgo.PermissionManageServer
	zero, _ = getFloatValue("0")
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
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "amount",
			Description: "Amount",
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
	amountStr := ""

	senderUserID := member.User.ID
	for _, o := range options {
		if o.Name == "user" {
			Var("Receiver User Option", o)
			receiverUserID = o.Value.(string)
		}

		if o.Name == "amount" {
			amountStr = o.StringValue()
		}
	}

	roleIDs := getRoleIDs(roles)
	message += getMessageDataFromUser(s, "Sender User", senderUserID, roleIDs)
	message += getMessageDataFromUser(s, "Receiver User", receiverUserID, roleIDs)
	message += fmt.Sprintln("Amount: ", amountStr)

	amountFloat, err := getFloatValue(amountStr)
	Var("Amount", amountFloat)
	if err != nil {
		message += fmt.Sprintln("Error parsing Amount: ", err)
	} else {
		if isLessThanZero(amountFloat) {
			message += fmt.Sprintln("Amount cannot be less than 0")
		}
	}

	printMessage(s, i, message, true)
}

func getFloatValue(value string) (*big.Float, error) {
	decimalValue, _, err := new(big.Float).Parse(value, 10)
	return decimalValue, err
}

func isLessThanZero(value *big.Float) bool {
	return value.Cmp(zero) < 1
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
