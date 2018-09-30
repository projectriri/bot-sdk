package qqbot

import (
	"errors"
	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	"github.com/projectriri/bot-sdk"
	"net/url"
)

type Bot struct {
	*qqbotapi.BotAPI
	WebhookConfig *qqbotapi.WebhookConfig
	self          User
}

func NewQQBot(token string, api string, secret string, updateConfig interface{}) (*Bot, error) {
	uc, ok := updateConfig.(qqbotapi.WebhookConfig)
	if !ok {
		return nil, errors.New("unsupported update config")
	}
	botapi, err := qqbotapi.NewBotAPI(token, api, secret)
	if err != nil {
		return nil, err
	}
	return &Bot{
		BotAPI:        botapi,
		WebhookConfig: &uc,
	}, nil
}

func (b *Bot) GetUpdatesChan(bufferSize int) (botsdk.UpdateChannel, botsdk.UpdateErrorChannel, error) {
	ch := b.BotAPI.ListenForWebhook(*b.WebhookConfig)
	updc := make(chan botsdk.Update)
	errc := make(chan error)
	go func() {
		for update := range ch {
			updc <- Update{
				Update: &update,
				bot:    b,
			}
		}
	}()
	return updc, errc, nil
}

func (b *Bot) Send(config botsdk.SendConfig, message interface{}) (messageID int64, err error) {
	var resp qqbotapi.Message
	switch m := message.(type) {
	case *url.URL:
		switch config.MediaType {
		case botsdk.TypeSticker:
			fallthrough
		case botsdk.TypeImage:
			msg := qqbotapi.NewImageWeb(m)
			if config.NoCache {
				msg.DisableCache()
			}
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		case botsdk.TypeRecord:
			msg := qqbotapi.NewRecordWeb(m)
			if config.NoCache {
				msg.DisableCache()
			}
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		}
	case url.URL:
		switch config.MediaType {
		case botsdk.TypeSticker:
			fallthrough
		case botsdk.TypeImage:
			msg := qqbotapi.NewImageWeb(&m)
			if config.NoCache {
				msg.DisableCache()
			}
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		case botsdk.TypeRecord:
			msg := qqbotapi.NewRecordWeb(&m)
			if config.NoCache {
				msg.DisableCache()
			}
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		}
	case string:
		switch config.MediaType {
		case botsdk.TypeSticker:
			fallthrough
		case botsdk.TypeImage:
			msg, _ := qqbotapi.NewImageBase64(m)
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		case botsdk.TypeRecord:
			msg, _ := qqbotapi.NewRecordBase64(m)
			resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, msg)
			messageID = resp.MessageID
		default: // TypeText
			if config.ParseMode == botsdk.FormatRawCQString {
				resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, m)
				messageID = resp.MessageID
			} else {
				resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, cqcode.Text{Text: m})
				messageID = resp.MessageID
			}
		}
	default:
		resp, err = b.BotAPI.SendMessage(config.ChatID, config.ChatType, message)
	}
	return
}

func (b *Bot) Self() (botsdk.User, error) {
	if b.self.ID != 0 {
		return b.self, nil
	}
	u, err := b.BotAPI.GetMe()
	if err != nil {
		return b.self, err
	}
	b.self.User = &u
	return b.self, err
}
