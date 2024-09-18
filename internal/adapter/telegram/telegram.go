package telegram

import (
	"errors"
	"log/slog"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Adapter struct {
}

var (
	ErrUnsupportedUpdate  = errors.New("unsupported update")
	ErrUnsupportedCommand = errors.New("unsupported command")
)

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) HandleUpdate(update *tgbotapi.Update) error {
	if update.Message == nil {
		return ErrUnsupportedUpdate
	}

	if update.Message.IsCommand() {
		return a.handleCommand(update.Message.Command())
	}
	return nil
}

func (a *Adapter) handleCommand(command string) error {
	switch command {
	case "distribution":
		slog.Info("distribution command")
	default:
		slog.Warn("unsupported command", "command", command)
		return ErrUnsupportedCommand
	}
	return nil
}
