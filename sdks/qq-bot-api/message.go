package qqbot

import (
	"github.com/catsworld/qq-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Message struct {
	*qqbotapi.Message
	bot *Bot
}

func (m Message) Bot() botsdk.Bot {
	return m.bot
}

func (m Message) MessageID() int64 {
	return int64(m.Message.MessageID)
}

func (m Message) Type() string {
	return "text"
}

func (m Message) Text() string {
	return m.Message.CQString()
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
