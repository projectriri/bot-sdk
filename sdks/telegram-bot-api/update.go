package tgbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Update struct {
	*tgbotapi.Update
	bot *Bot
}

func (u Update) UpdateID() int64 {
	return int64(u.Update.UpdateID)
}

func (u Update) Bot() botsdk.Bot {
	return u.bot
}

func (u Update) Message() botsdk.Message {
	var m *tgbotapi.Message
	if u.Update.Message != nil {
		m = u.Update.Message
	}
	if u.Update.EditedMessage != nil {
		m = u.Update.EditedMessage
	}
	if u.Update.ChannelPost != nil {
		m = u.Update.ChannelPost
	}
	if u.Update.EditedChannelPost != nil {
		m = u.Update.EditedChannelPost
	}
	return Message{
		Message: m,
		bot:     u.bot,
	}
}

func (u Update) IsMessage() bool {
	return u.Update.Message != nil ||
		u.Update.EditedMessage != nil ||
		u.Update.ChannelPost != nil ||
		u.Update.EditedChannelPost != nil
}
