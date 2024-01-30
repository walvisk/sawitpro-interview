package user

import (
	"context"
	"errors"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
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

	hashedPwd, err := utils.HashPassword(strings.TrimSpace(param.Password))
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

func (s *service) FindUserByPhone(c context.Context, phoneParam string) (*repository.User, error) {
	phone, countryCode := utils.GetPhoneAndCountryCode(strings.TrimSpace(phoneParam))

	user, err := s.repository.FindUserByPhoneAndCountryCode(c, phone, countryCode)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("wrong phone number")
	}

	return user, nil
}

func (s *service) FindUserByID(c context.Context, id int64) (*repository.User, error) {
	user, err := s.repository.FindUserByID(c, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
