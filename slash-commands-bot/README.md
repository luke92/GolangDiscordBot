# GolangDiscordBot
A sample Golang Discord Bot

# Tutorial
https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf

# Configure Application in Discord
https://discord.com/developers/applications

## Configure SCOPES
- Bot
- applications.commands

## Pre-requisites
Install Go in 1.16 version minimum.

# Install DiscordGo (a Client that interact with Go servers)
`go get github.com/bwmarrin/discordgo`

# Run bot
`$ go run main.go -guild {SERVERID} -token {BOT_TOKEN}`