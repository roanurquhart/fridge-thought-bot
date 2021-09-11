package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var submissionVotes map[string]int = make(map[string]int)
var suggestionVotes map[string]int = make(map[string]int)
var seqOrdered, regSeqOrdered = GenerateSequence()

func main() {
	fmt.Println("hello world")

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("No token provided. Please set env variable BOT_TOKEN")
		return
	}

	// Create a new Discord session using the provided bot token.
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	b.AddHandler(InputHandler)
	b.AddHandler(ReactionsHandler)

	// Open a websocket connection to Discord and begin listening.
	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to discord", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	b.Close()
}
