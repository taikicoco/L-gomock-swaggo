package repository

import (
	"context"
	"l-gomock/model"
)

type PlayerAPIRepository interface {
	GetPlayerList(context.Context) ([]*model.Player, error)
}
