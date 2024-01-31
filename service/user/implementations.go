package user

import (
	"context"
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
	phoneNumber, countryCode := utils.GetPhoneAndCountryCode(strings.TrimSpace(param.Phone))
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

	return user, nil
}

func (s *service) FindUserByID(c context.Context, id int64) (*repository.User, error) {
	user, err := s.repository.FindUserByID(c, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) UpdateUser(c context.Context, u *repository.User, fullName, phone string) error {
	var (
		phoneParam    string
		fullNameParam string
	)
	if phone == "" {
		phoneParam = u.Phone
	} else {
		phoneParam, _ = utils.GetPhoneAndCountryCode(phone)
	}

	if fullName == "" {
		fullNameParam = u.FullName
	} else {
		fullNameParam = strings.TrimSpace(fullName)
	}

	err := s.repository.UpdateUser(c, u, fullNameParam, phoneParam)
	if err != nil {
		return err
	}

	return nil
}
