package command

import "github.com/bwmarrin/discordgo"

var Leaderboard = []string{
	"3476e206-56b3-48e6-89bc-72e8f25906d3. New Idea. 2",
}

var LeaderboardCommand = discordgo.ApplicationCommand{
	Name:        "leaderboard",
	Description: "Show Project Leaderboard",
}

func LeaderboardCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Leaderboard[0],
		},
	})
}
