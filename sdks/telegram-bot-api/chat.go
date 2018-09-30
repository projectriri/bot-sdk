package tgbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectriri/bot-sdk"
)

type Chat struct {
	*tgbotapi.Chat
}

func (c Chat) ChatConfig() botsdk.ChatConfig {
	chatType := ""
	if c.IsPrivate() {
		chatType = "private"
	}
	if c.IsChannel() {
		chatType = "channel"
	}
	if c.IsGroup() {
		chatType = "group"
	}
	if c.IsSuperGroup() {
		chatType = "supergroup"
	}
	return botsdk.ChatConfig{
		ChatID:   c.ID,
		ChatType: chatType,
		Title:    c.Title,
	}
}
