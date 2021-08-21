package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var token string

func main() {
	fmt.Println("hello world")

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	b.AddHandler(inputHandler)
}
