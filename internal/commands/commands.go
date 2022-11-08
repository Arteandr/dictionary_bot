package commands

import tele "gopkg.in/telebot.v3"

type CMDHandler struct {
	Bot *tele.Bot
}

func NewCMDHandler(bot *tele.Bot) *CMDHandler {
	return &CMDHandler{
		Bot: bot,
	}
}
