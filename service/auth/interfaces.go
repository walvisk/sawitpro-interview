package auth

import (
	"github.com/SawitProRecruitment/UserService/repository"
)

type Service interface {
	AuthenticateUserPassword(*repository.User, string) error
	GenerateJWT() (string, error)
	ValidateJWT(string) error
}
