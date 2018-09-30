package qqbot

import (
	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	"github.com/projectriri/bot-sdk"
	"strconv"
)

type User struct {
	*qqbotapi.User
}

func (u User) Messenger() string {
	return "qq"
}

func (u User) UserID() int64 {
	return u.ID
}

func (u User) UserName() string {
	return ""
}

func (u User) FirstName() string {
	return u.User.NickName
}

func (u User) LastName() string {
	return ""
}

func (u User) DisplayName() string {
	if u.User.Card != "" {
		return u.User.Card
	}
	if u.User.NickName != "" {
		return u.User.NickName
	}
	return strconv.FormatInt(u.ID, 10)
}

func (u User) PrivateChat() botsdk.ChatConfig {
	return botsdk.ChatConfig{
		ChatID: u.ID,
	}
}

func (u User) At(parseMode int) string {
	switch parseMode {
	case botsdk.FormatRawCQString:
		return cqcode.FormatCQCode(&cqcode.At{
			QQ: strconv.FormatInt(u.ID, 10),
		})
	default:
		return "@" + u.DisplayName()
	}
}
