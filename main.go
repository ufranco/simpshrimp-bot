package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ufranco/simpshrimp-bot/bot"
	"github.com/ufranco/simpshrimp-bot/config"
)

func main() {
	config.ReadConfig()

	goBot, err := discordgo.New("Bot " + config.BotToken)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	bot.Start(goBot)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	goBot.Close()

}
