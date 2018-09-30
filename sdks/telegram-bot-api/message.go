package tgbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Message struct {
	*tgbotapi.Message
	bot *Bot
}

func (m Message) Bot() botsdk.Bot {
	return m.bot
}

func (m Message) MessageID() int64 {
	return int64(m.Message.MessageID)
}

func (m Message) Type() string {
	if m.Message.Text != "" {
		return "text"
	}
	if m.Sticker != nil {
		return "sticker"
	}
	if m.Photo != nil {
		return "photo"
	}
	return ""
}

func (m Message) Text() string {
	return m.Message.Text
}

func (m Message) Chat() botsdk.Chat {
	return Chat{
		m.Message.Chat,
	}
}

func (m Message) From() botsdk.User {
	return User{
		m.Message.From,
	}
}

func (m Message) Reply(config botsdk.MessageConfig, message interface{}) (interface{}, error) {
	return m.bot.Send(botsdk.SendConfig{
		MessageConfig: config,
		ChatConfig:    m.Chat().ChatConfig(),
	}, message)
}
