package main

import (
	"github.com/projectriri/bot-sdk/sdks/telegram-bot-api"
	"github.com/projectriri/bot-sdk/sdks/qq-bot-api"
	"github.com/projectriri/bot-sdk"
	"github.com/catsworld/qq-bot-api"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
)

func main()  {

	bots := make(botsdk.Bots, 0, 2)

	u1 := qqbotapi.NewWebhook("/webhook_endpoint")
	u1.PreloadUserInfo = true
	q, err := qqbot.NewQQBot("MyCoolqHttpToken", "http://localhost:5700", "CQHTTP_SECRET", u1)
	if err != nil {
		panic(err)
	}
	bots = append(bots, q)

	u2 := tgbotapi.NewUpdate(0)
	u2.Timeout = 60
	t, err := tgbot.NewTGBot("MyAwesomeBotToken", u2)
	if err != nil {
		panic(err)
	}
	bots = append(bots, t)

	ch, _, err := bots.GetUpdatesChan(100)
	if err != nil {
		panic(err)
	}

	for update := range ch {
		if update.IsMessage() {
			fmt.Println(update.Message().Text())
		}
	}

}