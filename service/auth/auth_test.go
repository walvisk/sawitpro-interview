package auth

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	gomock "github.com/golang/mock/gomock"
)

func TestAuthenticatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := NewMockService(ctrl)

	dummyPwd := "dolor1psum@"
	hashedPwd, err := utils.HashPassword(dummyPwd)
	if err != nil {
		t.FailNow()
	}

	tests := []struct {
		name    string
		payload *repository.User
		passwd  string
		returns error
		assert  func(error)
	}{
		{
			name: "when given correct password, returns no error",
			payload: &repository.User{
				ID:       int64(1),
				Password: hashedPwd,
			},
			passwd:  dummyPwd,
			returns: nil,
			assert: func(err error) {
				gomock.Eq(gomock.Nil()).Matches(err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authService.EXPECT().AuthenticateUserPassword(tt.payload, tt.passwd).Return(tt.returns)

			err := authService.AuthenticateUserPassword(tt.payload, tt.passwd)
			tt.assert(err)
		})
	}
}
