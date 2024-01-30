package user

import (
	"context"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository repository.RepositoryInterface
}

func NewUserService(repository repository.RepositoryInterface) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterUser(c context.Context, param generated.CreateUserJSONRequestBody) (int64, error) {
	countryCode, phoneNumber := param.Phone[:3], param.Phone[3:]
	u := repository.User{
		FullName:    strings.TrimSpace(param.FullName),
		CountryCode: strings.TrimSpace(countryCode),
		Phone:       strings.TrimSpace(phoneNumber),
	}

	hashedPwd, err := hashPassword(strings.TrimSpace(param.Password))
	if err != nil {
		return 0, err
	}
	u.Password = hashedPwd

	id, err := s.repository.CreateUser(c, &u)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
