package telegram

import (
	"context"
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type (
	Usecase interface {
	}

	API interface {
	}
)

type Bot struct {
	api *tgbotapi.BotAPI

	usecase Usecase
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api: api,
	}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	chatConfig := tgbotapi.NewUpdate(0)
	chatConfig.Timeout = 60

	chat, err := b.api.GetUpdatesChan(chatConfig)
	if err != nil {
		return err
	}

	for message := range chat {
		if message.Message == nil {
			continue
		}

		fmt.Println(message.Message)
	}

	return nil
}
