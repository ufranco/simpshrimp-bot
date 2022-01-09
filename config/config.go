package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	BotToken  string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	BotToken  string `json:"BOT_TOKEN"`
	BotPrefix string `json:"BOT_PREFIX"`
}

func ReadConfig() {
	log.Printf("Loading env variables...")

	BotToken = os.Getenv("BOT_TOKEN")

	if BotToken != "" {
		return
	}

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	BotToken = config.BotToken
}
