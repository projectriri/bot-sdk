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
		Messenger: "telegram",
		ChatID:    c.ID,
		ChatType:  chatType,
		Title:     c.Title,
	}
}

type User struct {
	*tgbotapi.User
}

func (u User) Messenger() string {
	return "telegram"
}

func (u User) UserID() int64 {
	return int64(u.ID)
}

func (u User) UserName() string {
	return u.User.UserName
}

func (u User) FirstName() string {
	return u.User.FirstName
}

func (u User) LastName() string {
	return u.User.LastName
}

func (u User) DisplayName() string {
	return u.User.FirstName
}

func (u User) PrivateChat() botsdk.ChatConfig {
	return botsdk.ChatConfig{
		Messenger: "telegram",
		ChatID:    int64(u.ID),
	}
}

type Message struct {
	*tgbotapi.Message
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

type Update struct {
	*tgbotapi.Update
	bot *Bot
}

func (u Update) UpdateID() int64 {
	return int64(u.Update.UpdateID)
}
