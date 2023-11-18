package repository

import (
	"l-gomock-swaggo/model"
)

type User interface {
	Update(user *model.User) error
	Get(id int) (*model.User, error)
}
