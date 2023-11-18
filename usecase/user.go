package usecase

import (
	"l-gomock-swaggo/model"
	"l-gomock-swaggo/repository"
)

type User struct {
	repo repository.User
}

func (u *User) Update(user *model.User) error {
	return u.repo.Update(user)
}

func (u *User) Get(id int) (*model.User, error) {
	return u.repo.Get(id)
}
