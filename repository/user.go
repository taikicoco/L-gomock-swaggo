package repository

import (
	"l-gomock/model"
)

type User interface {
	Update(user *model.User) (int, error)
}
