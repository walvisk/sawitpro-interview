package userLog

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository"
)

type Service interface {
	CreateUserLog(c context.Context, u *repository.User) error
}
