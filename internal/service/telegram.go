package service

import (
	"context"

	"github.com/AbrakadabraK/telegram-bot-pet/model"
)

type TelegramManager interface {
	GetUpdates(ctx context.Context, offset int, limit int) ([]*model.Update, error)
}
