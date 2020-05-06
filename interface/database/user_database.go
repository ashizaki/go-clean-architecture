package database

import (
	"context"

	"github.com/ashizaki/go-clean-architecture/domain/model"
	"github.com/ashizaki/go-clean-architecture/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

// ErrorMsg generates and returns error message.
func (r *userRepository) ErrorMsg(method model.RepositoryMethod, err error) error {
	return &model.RepositoryError{
		BaseErr:          err,
		RepositoryMethod: method,
		DomainModelName:  model.DomainModelNameUser,
	}
}

func (r *userRepository) Store(_ context.Context, handler repository.SqlHandler, u *model.User) (uint, error) {
	if err := handler.Create(u).Error(); err != nil {
		err = r.ErrorMsg(model.RepositoryMethodInsert, err)
		return 0, err
	}
	return u.ID, nil
}

func (r *userRepository) Fetch(_ context.Context, handler repository.SqlHandler, limit int, cursor uint) (*model.UserList, error) {
	var users []model.User
	if err := handler.Limit(limitForHasNext(limit)).Offset(cursorToOffset(cursor)).Select(&users).Error(); err != nil {
		err = r.ErrorMsg(model.RepositoryMethodList, err)
		return nil, err
	}

	if checkHasNext(len(users), limit) {
		return &model.UserList{
			Users:   users[:limit],
			HasNext: true,
			Cursor:  cursor + uint(limit),
		}, nil
	} else {
		return &model.UserList{
			Users:   users,
			HasNext: false,
			Cursor:  0,
		}, nil
	}
}

func (r *userRepository) GetById(_ context.Context, handler repository.SqlHandler, id uint) (*model.User, error) {
	user := &model.User{}
	if err := handler.Get(user, "id = ?", id).Error(); err != nil {
		if handler.IsRecordNotFoundError(err) {
			return nil, nil
		}
		err = r.ErrorMsg(model.RepositoryMethodRead, err)
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByAccount(_ context.Context, handler repository.SqlHandler, account string) (*model.User, error) {
	user := &model.User{}
	if err := handler.Get(user, "account = ?", account).Error(); err != nil {
		if handler.IsRecordNotFoundError(err) {
			return nil, nil
		}
		err = r.ErrorMsg(model.RepositoryMethodRead, err)
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(_ context.Context, handler repository.SqlHandler, id uint, u *model.User) error {
	u.ID = id
	if err := handler.Save(u).Error(); err != nil {
		err = r.ErrorMsg(model.RepositoryMethodUpdate, err)
		return err
	}
	return nil
}

func (r *userRepository) Delete(_ context.Context, handler repository.SqlHandler, id uint) error {
	user := &model.User{}
	if err := handler.Delete(user, "id = ?", id).Error(); err != nil {
		err = r.ErrorMsg(model.RepositoryMethodDelete, err)
		return err
	}
	return nil
}
