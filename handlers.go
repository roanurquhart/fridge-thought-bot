package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func InputHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the !checkers prefix
	if !strings.HasPrefix(m.Content, "!fridge") {
		return
	}

	// Get the arguments
	args := strings.Split(m.Content, " ")[1:]
	// Ensure valid command
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Command missing: For a list of commands type !fridge help")
		return
	}

	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	switch args[0] {
	case "help":
		title = "Fridge Help"
		description = "Pick a topic below to get help"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Suggest",
				Value: "`!fridge suggest <your phrase>`:  Command to suggest a phrase based on the prompt",
			},
			{
				Name:  "Vote",
				Value: "Add your reaction to suggestion to vote",
			},
			{
				Name:  "Check",
				Value: "Check today's letters",
			},
			{
				Name:  "Submit",
				Value: "!fridge submit <your phrase>: Command to submit a phrase using the letters of the day",
			},
		}
	case "check":
		s.ChannelMessageSend(m.ChannelID, "Letters for the day: "+seqOrdered)
		// s.Channel()
	case "suggest":
		suggestionHandler(s, m, strings.Join(args[2:], " "))
	case "submit":
		fmt.Println(strings.Join(args[1:], " "))
		submissionHandler(s, m, strings.Join(args[1:], " "), regSeqOrdered)
		fmt.Printf("Channel ID %s", string(m.ChannelID))
	default:
		s.ChannelMessageSend(m.ChannelID, "Invalid command. For a list of help topics, type !checkers help")

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
	})
}

func submissionHandler(s *discordgo.Session, m *discordgo.MessageCreate, phrase string, seqReg string) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Bot error. Error getting channel.")
		return
	}

	// Ensure that the command is not being sent from a dm
	if c.Type == discordgo.ChannelTypeDM {
		s.ChannelMessageSend(m.ChannelID, "Invalid channel. Cannot send invites from a DM")
		return
	}

	if InSequence(seqReg, phrase) != true {
		s.ChannelMessageSend(m.ChannelID, "Phrase "+phrase+" is Wrong letters!")
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Phrase "+phrase+" is Right letters!")

	if _, ok := submissionVotes[phrase]; ok {
		s.ChannelMessageSend(m.ChannelID, "Submission has already been submitted")
		return
	} else {
		submissionVotes[phrase] = 0
	}

	submission, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Fridge Thought submission  %s", phrase),
		Description: fmt.Sprintf("From: %s", m.Author.Username),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s:%s", "submission", phrase),
		},
	})
	if err != nil {
		log.Panicf("Bot was unable to send message to channel with ID: %s", m.ChannelID)
	}
	s.MessageReactionAdd(m.ChannelID, submission.ID, "❤️")
	s.MessageReactionAdd(m.ChannelID, submission.ID, "💩")

}

func suggestionHandler(s *discordgo.Session, m *discordgo.MessageCreate, phrase string) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Bot error. Error getting channel.")
		return
	}

	// Ensure that the command is not being sent from a dm
	if c.Type == discordgo.ChannelTypeDM {
		s.ChannelMessageSend(m.ChannelID, "Invalid channel. Cannot send invites from a DM")
		return
	}

	if _, ok := suggestionVotes[phrase]; ok {
		s.ChannelMessageSend(m.ChannelID, "Suggestion has already been submitted")
		return
	} else {
		suggestionVotes[phrase] = 0
	}

	suggestion, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Fridge Thought suggestion  %s", phrase),
		Description: fmt.Sprintf("From: %s", m.Author.Username),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s:%s", "suggestion", phrase),
		},
	})
	if err != nil {
		log.Panicf("Bot was unable to send message to channel with ID: %s", m.ChannelID)
	}
	s.MessageReactionAdd(m.ChannelID, suggestion.ID, "❤️")
	s.MessageReactionAdd(m.ChannelID, suggestion.ID, "💩")
}

func ReactionsHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	// Fetch some extra information about the message associated to the reaction
	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	// Ignore reactions on messages that have an error or that have not been sent by the bot
	if err != nil || m == nil || m.Author.ID != s.State.User.ID {
		return
	}

	// Ignore messages that are not embeds with a command in the footer
	if len(m.Embeds) != 1 || m.Embeds[0].Footer == nil || m.Embeds[0].Footer.Text == "" {
		return
	}

	// Ignore reactions that haven't been set by the bot
	if !isBotReaction(s, m.Reactions, &r.Emoji) {
		return
	}

	user, err := s.User(r.UserID)
	// Ignore when sender is invalid or is a bot
	if err != nil || user == nil || user.Bot {
		return
	}

	args := strings.Split(m.Embeds[0].Footer.Text, ":")
	// Ensure valid footer command
	if len(args) != 2 {
		return
	}

	// Call the corresponding handler
	switch args[0] {
	case "suggestion":
		suggestionReactionHandler(s, r, m, args[1])
	case "submission":
		submissionReactionHandler(s, r, m, args[1])
	}

}

func suggestionReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, phrase string) {
	if r.Emoji.Name == "❤️" {
		suggestionVotes[phrase] += 1
	}
	if r.Emoji.Name == "💩" {
		suggestionVotes[phrase] -= 1
	}
}

func submissionReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, phrase string) {
	if r.Emoji.Name == "❤️" {
		submissionVotes[phrase] += 1
	}
	if r.Emoji.Name == "💩" {
		submissionVotes[phrase] -= 1
	}
}
