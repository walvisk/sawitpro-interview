package user

import (
	"context"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	gomock "github.com/golang/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	useServiceMock := NewMockService(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name    string
		payload generated.CreateUserJSONRequestBody
		returns []any
		assert  func(int64, error)
	}{
		{
			name: "when given correct payload, returns id and no error",
			payload: generated.CreateUserJSONRequestBody{
				FullName: "Dolor Ipsum",
				Phone:    "+628132467662",
				Password: "d0lorIpsum@",
			},
			returns: []any{int64(1), nil},
			assert: func(i int64, err error) {
				gomock.Eq(int64(1)).Matches(i)
				gomock.Eq(gomock.Nil()).Matches(err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useServiceMock.EXPECT().RegisterUser(ctx, tt.payload).Return(tt.returns...)

			id, err := useServiceMock.RegisterUser(ctx, tt.payload)
			tt.assert(id, err)
		})
	}
}
