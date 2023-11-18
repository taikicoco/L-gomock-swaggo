package usecase

import (
	"l-gomock-swaggo/model"
	"l-gomock-swaggo/repository/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mock.NewMockUser(ctrl)
	mockUser.EXPECT().Update(&model.User{ID: 1, Name: "test"}).Return(nil)

	user := &User{repo: mockUser}
	user.repo = mockUser
	u := &model.User{ID: 1, Name: "test"}

	err := user.Update(u)
	if err != nil {
		t.Errorf("failed test %#v", err)
	}
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mock.NewMockUser(ctrl)
	mockUser.EXPECT().Get(1).Return(&model.User{ID: 1, Name: "test"}, nil)

	user := &User{repo: mockUser}
	user.repo = mockUser
	expected := &model.User{ID: 1, Name: "test"}

	s, err := user.Get(1)
	if err != nil {
		t.Errorf("failed test %#v", err)
	}
	if s.ID != expected.ID {
		t.Errorf("failed test %#v", s.ID)
	}
	if s.Name != expected.Name {
		t.Errorf("failed test %#v", s.Name)
	}
}
