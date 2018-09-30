package qqbot

import (
	"github.com/catsworld/qq-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Chat struct {
	*qqbotapi.Chat
}

func (c Chat) ChatConfig() botsdk.ChatConfig {
	return botsdk.ChatConfig{
		ChatID:   c.ID,
		ChatType: c.Type,
		// TODO: Title
	}
}
