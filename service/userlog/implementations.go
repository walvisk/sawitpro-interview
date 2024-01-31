package userlog

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository"
)

type service struct {
	repository repository.RepositoryInterface
}

func NewUserLogService(repository repository.RepositoryInterface) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateUserLog(c context.Context, u *repository.User) error {
	err := s.repository.CreateUserLog(c, u)
	if err != nil {
		return err
	}

	return nil
}
