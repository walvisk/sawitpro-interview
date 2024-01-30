package user

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Service interface {
	RegisterUser(context.Context, generated.CreateUserJSONRequestBody) (int64, error)
	FindUserByPhone(context.Context, string) (*repository.User, error)
	FindUserByID(context.Context, int64) (*repository.User, error)
}
