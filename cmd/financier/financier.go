package main

import (
	"context"
	telegramAdapter "financier/internal/adapter/telegram"
	"financier/internal/config"
	"financier/internal/driver/telegram"
	"os"

	"log/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	telegramBotAdapter := telegramAdapter.NewAdapter()

	telegramBot, err := telegram.NewBot(cfg.TelegamToken, telegramBotAdapter)
	if err != nil {
		slog.Error("failed to create telegram bot", "err", err)
		os.Exit(1)
	}

	ctx := context.Background()
	if err := telegramBot.Run(ctx); err != nil {
		slog.Error("failed to run telegram bot", "err", err)
		os.Exit(1)
	}
}
