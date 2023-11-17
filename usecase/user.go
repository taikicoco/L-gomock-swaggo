package usecase

import (
	"l-gomock/model"
	"l-gomock/repository"
)

type User struct{
	repo repository.User
}

func (u *User) Update(user *model.User) (int, error) {
	return u.repo.Update(user)
}
