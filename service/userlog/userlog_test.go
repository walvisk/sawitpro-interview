package userlog

import (
	"context"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	gomock "github.com/golang/mock/gomock"
)

func TestCreateUserLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userLogServiceMock := NewMockService(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name    string
		payload *repository.User
		returns error
		assert  func(error)
	}{
		{
			name: "when given valid payload, create data correctly",
			payload: &repository.User{
				ID:       int64(1),
				FullName: "Dolor Ipsum",
			},
			returns: nil,
			assert: func(err error) {
				gomock.Eq(gomock.Nil()).Matches(err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userLogServiceMock.EXPECT().CreateUserLog(ctx, tt.payload).Return(tt.returns)

			err := userLogServiceMock.CreateUserLog(ctx, tt.payload)
			tt.assert(err)
		})
	}
}
