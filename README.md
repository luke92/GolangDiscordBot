# GolangDiscordBot
A sample Golang Discord Bot

# Tutorial
https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf

# Configure Application in Discord
https://discord.com/developers/applications

## Pre-requisites
Install Go in 1.16 version minimum.

# Install DiscordGo (a Client that interact with Go servers)
`go get github.com/bwmarrin/discordgo`

## Build the app

`$ go build -o bin/modak-test-vote-bot-discord main.go`

or

`$ task build`

## Run the app

First, you can export the token:
`$ export BOT_TOKEN=<your bot token>`

Second, expor the API
`$ export API_URL=<your base api url>`

Let's run locally our Bot right now:

```
$ ./bin/modak-test-vote-bot-discord -t $BOT_TOKEN -api $API_URL
Bot is now running.  Press CTRL-C to exit.
```

or

`$ task bot`

## Test the app

```
$ task bot
task: [bot] ./bin/modak-test-vote-bot-discord -t $BOT_TOKEN -api $API_URL
Bot is now running.  Press CTRL-C to exit.
```

Now you can tape `/vote`, `/project` and `/projectleaderboard` in the Discord server connected to the Bot :).