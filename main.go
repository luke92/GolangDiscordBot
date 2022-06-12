package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
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
	Token      string
	API_URL    string
	PREFIX     string
	Categories []string
	Projects   []string
)

//const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
//const API_URL = https://9f801b5b-d44a-497f-8405-3b7d687a0cd3.mock.pstmn.io

func init() {
	PREFIX = "!"
	Categories = []string{
		"CODING",
		"ART",
		"OTHER",
		"FOOD",
		"FASHION",
	}
	Projects = []string{
		"3476e206-56b3-48e6-89bc-72e8f25906d3. New Idea. Description of idea",
	}
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

	if strings.HasPrefix(m.Content, PREFIX+"vote") {
		voteCommand(s, m)
	}

	if strings.HasPrefix(m.Content, PREFIX+"project") {
		projectCommand(s, m)
	}

	if strings.HasPrefix(m.Content, PREFIX+"leaderboard") {
		projectLeaderboard(s, m)
	}
}

func voteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	values := getValuesFromCommand(m.Content)
	message := ""
	if len(values) == 1 {
		message = "You are not enter a Project ID"
	} else if !IsValidUUID(values[1]) {
		message = "You are not enter a valid Project ID"
	} else {
		message = "You are vote the project " + values[1] + ".The new total amount of votes for the project is 1."
	}
	sendMessageTextToChannel(s, m, message)
}

func projectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	values := getValuesFromCommand(m.Content)
	message := Projects[0]
	if len(values) > 1 {
		category := strings.ToUpper(values[1])
		if !contains(Categories, category) {
			message = "You not enter a valid category"
		}
	}
	sendMessageTextToChannel(s, m, message)
}

func projectLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendMessageTextToChannel(s, m, Projects[0])
}

func sendMessageTextToChannel(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		fmt.Println(err)
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func getValuesFromCommand(content string) []string {
	inputMessage := strings.TrimSpace(content)
	values := strings.Split(inputMessage, " ")
	return values
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
