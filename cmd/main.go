package main

import (
	"dictionary_bot/internal/commands"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatalf("[ERROR] Bot token is missing")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("[ERROR] Bot starting error: %s", err.Error())
	}

	cmdHandler := commands.NewCMDHandler(bot)

	bot.Handle(tele.OnText, cmdHandler.OnText)

	bot.Start()
}
