package telegram

import (
	"context"
	"log/slog"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type (
	Usecase interface {
	}

	Adapter interface {
		HandleUpdate(update *tgbotapi.Update) error
	}
)

type Bot struct {
	api *tgbotapi.BotAPI

	adapter Adapter
}

func NewBot(token string, adapter Adapter) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:     api,
		adapter: adapter,
	}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	chatConfig := tgbotapi.NewUpdate(0)
	chatConfig.Timeout = 60

	chat, err := b.api.GetUpdatesChan(chatConfig)
	if err != nil {
		return err
	}

	for updateEvent := range chat {
		if err := b.adapter.HandleUpdate(&updateEvent); err != nil {
			slog.Error("failed to handle telegram update", "err", err)
		}
	}

	return nil
}
