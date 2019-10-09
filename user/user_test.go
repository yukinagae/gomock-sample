package user

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	"github.com/golang/mock/gomock"
	mock_user "github.com/yukinagae/gomock-sample/user/mock_user"
)

func TestRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_user.NewMockRepository(ctrl)

	mockRepository.EXPECT().Get(1).Return(&User{1, "hoge", "hoge@example.com"})

	var repo user.Repository
	repo = mockRepository
	user := repo.Get(1)
	assert.Equal(t, user.ID, 1)
}
