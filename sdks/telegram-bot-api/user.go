package tgbot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectriri/bot-sdk"
	"strings"
)

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
		ChatID: int64(u.ID),
	}
}

func (u User) At(parseMode int) string {
	switch parseMode {
	case botsdk.FormatMarkdown:
		if u.User.UserName != "" {
			return "@" + strings.Replace(u.User.UserName, "_", "\\_", -1)
		} else {
			return fmt.Sprintf("[%s](tg://%s)", u.User.FirstName, u.User.ID)
		}
	case botsdk.FormatHTML:
		if u.User.UserName != "" {
			return "@" + u.User.UserName
		} else {
			return fmt.Sprintf("<a href='tg://%s'>%s</a>", u.User.ID, u.User.FirstName)
		}
	default:
		if u.User.UserName != "" {
			return "@" + u.User.UserName
		} else {
			return "@" + u.User.FirstName
		}
	}
}
