package bot

import (
	"log"
	"os"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func Start(goBot *discordgo.Session) {

	user, err := goBot.User("@me")

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	BotID := user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	log.Printf("Bot running! Id: %s. Press Ctrl-C to exit", BotID)

}

func messageHandler(discordSession *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discordSession.State.User.ID || message.GuildID == "" {
		return
	}

	jokePattern, err := regexp.MatchString("(?i)^(11|once|(.*\\s(11|once)))(\\.|\\.\\.\\.|\\?|!){0,1}$", message.Content)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if jokePattern {
		log.Printf("%s: %s", message.Author.Username, message.Content)
		discordSession.ChannelMessageSend(message.ChannelID, "chupalo entonce XD")
	}

}
