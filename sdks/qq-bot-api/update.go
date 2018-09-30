package qqbot

import (
	"github.com/catsworld/qq-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Update struct {
	*qqbotapi.Update
	bot *Bot
}

func (u Update) UpdateID() int64 {
	return int64(u.Update.MessageID)
}

func (u Update) Bot() botsdk.Bot {
	return u.bot
}

func (u Update) Message() botsdk.Message {
	return Message{
		Message: u.Update.Message,
		bot:     u.bot,
	}
}

func (u Update) IsMessage() bool {
	return u.Update.Message != nil
}
