package user

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Service interface {
	RegisterUser(c context.Context, params generated.CreateUserJSONRequestBody) (int64, error)
	FindUserByPhone(c context.Context, phone string) (*repository.User, error)
	FindUserByID(c context.Context, id int64) (*repository.User, error)
	UpdateUser(c context.Context, u *repository.User, phone, fullName string) error
}
