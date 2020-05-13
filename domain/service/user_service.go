package service

import (
	"context"

	"github.com/ashizaki/go-clean-architecture/domain/repository"
	"github.com/pkg/errors"
)

type UserService interface {
	IsAlreadyExistID(ctx context.Context, handler repository.SqlHandler, id uint) (bool, error)
	IsAlreadyExistAccount(ctx context.Context, handler repository.SqlHandler, account string) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) IsAlreadyExistID(ctx context.Context, handler repository.SqlHandler, id uint) (bool, error) {
	user, err := s.repo.GetById(ctx, handler, id)
	if err != nil {
		return false, errors.Wrap(err, "failed to get user by id")
	}
	return user != nil, nil
}

func (s *userService) IsAlreadyExistAccount(ctx context.Context, handler repository.SqlHandler, account string) (bool, error) {
	user, err := s.repo.GetByAccount(ctx, handler, account)
	if err != nil {
		return false, errors.Wrap(err, "failed to get user by account")
	}
	return user != nil, nil
}
