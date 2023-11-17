package usecase

import (
	"l-gomock/model"
	"l-gomock/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mock.NewMockUser(ctrl)
	mockUser.EXPECT().Update(&model.User{ID: 1, Name: "test"}).Return(1, nil)

	user := &User{repo: mockUser}
	user.repo = mockUser
	expected := &model.User{ID: 1, Name: "test"}

	s, err := user.Update(expected)
	if err != nil {
		t.Errorf("failed test %#v", err)
	}
	if s != 1 {
		t.Errorf("failed test %#v", s)
	}
}
