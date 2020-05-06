package repository

import (
	"context"

	"github.com/ashizaki/go-clean-architecture/domain/model"
)

type UserRepository interface {
	Store(ctx context.Context, handler SqlHandler, u *model.User) (uint, error)
	Fetch(ctx context.Context, handler SqlHandler, limit int, cursor uint) (*model.UserList, error)
	GetById(ctx context.Context, handler SqlHandler, id uint) (*model.User, error)
	GetByAccount(ctx context.Context, handler SqlHandler, account string) (*model.User, error)
	Update(ctx context.Context, handler SqlHandler, id uint, u *model.User) error
	Delete(ctx context.Context, handler SqlHandler, id uint) error
}
