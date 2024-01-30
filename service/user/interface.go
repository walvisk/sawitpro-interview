package user

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
)

type Service interface {
	RegisterUser(context.Context, generated.CreateUserJSONRequestBody) (int64, error)
}
