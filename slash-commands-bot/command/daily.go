package command

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/bwmarrin/discordgo"
)

var DailyCommand = discordgo.ApplicationCommand{
	Name:        "daily",
	Description: "Copy of Daily",
}

func DailyCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	min := int64(1)
	max := int64(10)

	randNumber, _ := rand.Int(rand.Reader, big.NewInt(max))
	amount := randNumber.Int64() + min

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Color: 39423,
					Author: &discordgo.MessageEmbedAuthor{
						Name:    i.Member.User.String(),
						IconURL: i.Member.User.AvatarURL(""),
					},
					Footer: &discordgo.MessageEmbedFooter{
						IconURL: "https://i.imgur.com/AfFp7pu.png",
						Text:    "Daily Footer",
					},
					Title: "Daily Reward",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Amount",
							Value:  fmt.Sprintf("%d points", amount),
							Inline: true,
						},
						{
							Name:   "Next reward",
							Value:  "in 1d",
							Inline: true,
						},
					},
				},
			},
			Content: "",
		},
	})
}
