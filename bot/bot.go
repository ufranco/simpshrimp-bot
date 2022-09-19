package bot

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var jokePattern = regexp.MustCompile("(11|once|(.*\\s(11|once)))(\\.|\\.\\.\\.|\\?|!){0,1}$")

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
	fmt.Println(message.Author)
	if message.Author.ID == discordSession.State.User.ID || message.GuildID == "" {
		return
	}

	isJokeable := jokePattern.MatchString(message.Message.Content)

	if isJokeable {
		discordSession.ChannelMessageSend(message.ChannelID, "chupalo entonce XD")
	}

}
