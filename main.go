package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

/*
type ProjectData struct {
	Project_UUID        string `json: "project_uuid"`
	Project_Title       string `json: "project_title"`
	Project_Total_Votes int    `json: "project_total_votes"`
}

type VoteResponse struct {
	Message string `json: "message"`
}
*/

// Variables used for command line parameters
var (
	Token   string
	API_URL string
)

//const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
//const API_URL = https://9f801b5b-d44a-497f-8405-3b7d687a0cd3.mock.pstmn.io

func init() {
	//Receive parameters from arguments
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&API_URL, "api", "", "BASE API URL")
	flag.Parse()
	fmt.Println(Token)
	fmt.Println(API_URL)
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := createDiscordSession()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	setupIntents(dg)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	runBot()

	// Cleanly close down the Discord session.
	dg.Close()
}

func createDiscordSession() (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + Token)
	return dg, err
}

func setupIntents(dg *discordgo.Session) {
	dg.Identify.Intents = discordgo.IntentsGuildMessages
}

func runBot() {
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!vote" {
		voteCommand(s, m)
	}

	if m.Content == "!project" {
		projectCommand(s, m)
	}

	if m.Content == "!projectleaderboard" {
		projectLeaderboard(s, m)
	}
}

func voteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendMessageTextToChannel(s, m, "You are voting")
}

func projectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendMessageTextToChannel(s, m, "You are getting values of Project")
}

func projectLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendMessageTextToChannel(s, m, "Leaderboard")
}

func sendMessageTextToChannel(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		fmt.Println(err)
	}
}
