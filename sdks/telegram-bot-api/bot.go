package tgbot

import (
	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectriri/bot-sdk"
	"net/url"
)

type Bot struct {
	*tgbotapi.BotAPI
	UpdateConfig *tgbotapi.UpdateConfig
	self         User
}

func NewTGBot(token string, updateConfig interface{}) (*Bot, error) {
	uc, ok := updateConfig.(tgbotapi.UpdateConfig)
	if !ok {
		return nil, errors.New("unsupported update config")
	}
	botapi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		BotAPI:       botapi,
		UpdateConfig: &uc,
	}, nil
}

func (b *Bot) GetUpdatesChan(bufferSize int) (botsdk.UpdateChannel, botsdk.UpdateErrorChannel, error) {
	ch, err := b.BotAPI.GetUpdatesChan(*b.UpdateConfig)
	if err != nil {
		return nil, nil, err
	}
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

func (b *Bot) sendFile(config botsdk.SendConfig, m interface{}) (messageID int64, err error) {
	var resp tgbotapi.Message
	switch config.MediaType {
	case botsdk.TypeSticker:
		msg := tgbotapi.NewStickerUpload(config.ChatID, m)
		msg.ReplyToMessageID = int(config.ReplyToMessageID)
		resp, err = b.BotAPI.Send(msg)
		messageID = int64(resp.MessageID)
	case botsdk.TypeImage:
		msg := tgbotapi.NewPhotoUpload(config.ChatID, m)
		msg.ReplyToMessageID = int(config.ReplyToMessageID)
		resp, err = b.BotAPI.Send(msg)
		messageID = int64(resp.MessageID)
	case botsdk.TypeRecord:
		msg := tgbotapi.NewVoiceUpload(config.ChatID, m)
		msg.ReplyToMessageID = int(config.ReplyToMessageID)
		resp, err = b.BotAPI.Send(msg)
		messageID = int64(resp.MessageID)
	}
	return
}

func (b *Bot) Send(config botsdk.SendConfig, message interface{}) (messageID int64, err error) {
	if config.MediaType == botsdk.TypeDeleteMessage {
		msg := tgbotapi.DeleteMessageConfig{
			ChatID:    config.ChatID,
			MessageID: int(config.MessageID),
		}
		resp, err := b.BotAPI.Send(msg)
		return int64(resp.MessageID), err
	}
	var resp tgbotapi.Message
	switch m := message.(type) {
	case tgbotapi.Chattable:
		resp, err = b.BotAPI.Send(m)
		messageID = int64(resp.MessageID)
	case *url.URL:
		return b.sendFile(config, *m)
	case url.URL:
		return b.sendFile(config, m)
	case tgbotapi.FileBytes:
		return b.sendFile(config, m)
	case tgbotapi.FileReader:
		return b.sendFile(config, m)
	default:
		str, ok := m.(string)
		if !ok {
			str = fmt.Sprint(m)
		}
		switch config.MediaType {
		case botsdk.TypeSticker:
			msg := tgbotapi.NewStickerShare(config.ChatID, str)
			msg.ReplyToMessageID = int(config.ReplyToMessageID)
			resp, err = b.BotAPI.Send(msg)
			messageID = int64(resp.MessageID)
		case botsdk.TypeImage:
			msg := tgbotapi.NewPhotoShare(config.ChatID, str)
			msg.ReplyToMessageID = int(config.ReplyToMessageID)
			resp, err = b.BotAPI.Send(msg)
			messageID = int64(resp.MessageID)
		case botsdk.TypeRecord:
			msg := tgbotapi.NewVoiceShare(config.ChatID, str)
			msg.ReplyToMessageID = int(config.ReplyToMessageID)
			resp, err = b.BotAPI.Send(msg)
			messageID = int64(resp.MessageID)
		case botsdk.TypeEditMessage:
			msg := tgbotapi.NewEditMessageText(config.ChatID, int(config.MessageID), str)
			resp, err = b.BotAPI.Send(msg)
			messageID = int64(resp.MessageID)
		default: // TypeText
			msg := tgbotapi.NewMessage(config.ChatID, str)
			switch config.ParseMode {
			case botsdk.FormatHTML:
				msg.ParseMode = "HTML"
			case botsdk.FormatMarkdown:
				msg.ParseMode = "Markdown"
			}
			msg.DisableWebPagePreview = config.DisableWebPreview
			msg.ReplyToMessageID = int(config.ReplyToMessageID)
			resp, err = b.BotAPI.Send(msg)
			messageID = int64(resp.MessageID)
		}
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
