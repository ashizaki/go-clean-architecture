package interactor

import (
	"context"

	"github.com/ashizaki/go-clean-architecture/domain/model"
	"github.com/ashizaki/go-clean-architecture/domain/repository"
	"github.com/ashizaki/go-clean-architecture/domain/service"
	"github.com/pkg/errors"
)

type UserInteractor interface {
	Store(ctx context.Context, u *model.User) (*model.User, error)
	Fetch(ctx context.Context, limit int, cursor uint) (res *model.UserList, err error)
	Get(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, id uint, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
}

type userInteractor struct {
	db      repository.DbHandler
	service service.UserService
	repo    repository.UserRepository
}

func NewUserInteractor(db repository.DbHandler, service service.UserService, repo repository.UserRepository) UserInteractor {
	return &userInteractor{
		db:      db,
		service: service,
		repo:    repo,
	}
}

func (i *userInteractor) Store(ctx context.Context, u *model.User) (*model.User, error) {
	txHandler, err := i.db.Begin()
	if err != nil {
		return nil, beginTxErrorMsg(err)
	}

	defer func() {
		if err := txHandler.End(err); err != nil {
			err = errors.Wrap(err, "failed to close tx")
		}
	}()

	exist, err := i.service.IsAlreadyExistAccount(ctx, txHandler, u.Account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to is already exist Account")
	}

	if exist {
		err = &model.AlreadyExistError{
			PropertyName:    model.AccountProperty,
			PropertyValue:   u.Account,
			DomainModelName: model.DomainModelNameUser,
		}
		return nil, errors.Wrap(err, "already exist id")
	}

	id, err := i.repo.Store(ctx, txHandler, u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert user")
	}

	u.ID = id
	return u, err
}

func (i *userInteractor) Fetch(ctx context.Context, limit int, cursor uint) (res *model.UserList, err error) {
	list, err := i.repo.Fetch(ctx, i.db, limit, cursor)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	return list, nil
}

func (i *userInteractor) Get(ctx context.Context, id uint) (*model.User, error) {
	user, err := i.repo.GetById(ctx, i.db, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by id")
	}
	return user, nil
}

func (i *userInteractor) Update(ctx context.Context, id uint, u *model.User) (user *model.User, err error) {
	txHandler, err := i.db.Begin()
	if err != nil {
		return nil, beginTxErrorMsg(err)
	}

	defer func() {
		if err := txHandler.End(err); err != nil {
			err = errors.Wrap(err, "failed to close tx")
		}
	}()

	exist, err := i.service.IsAlreadyExistID(ctx, txHandler, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to is already exist ID")
	}

	if !exist {
		err = &model.NoSuchDataError{
			PropertyName:    model.IDProperty,
			PropertyValue:   id,
			DomainModelName: model.DomainModelNameUser,
		}
		return nil, errors.Wrap(err, "does not exists ID")
	}

	copied := *u

	if err := i.repo.Update(ctx, txHandler, id, &copied); err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return &copied, nil
}

func (i *userInteractor) Delete(ctx context.Context, id uint) error {
	txHandler, err := i.db.Begin()
	if err != nil {
		return beginTxErrorMsg(err)
	}

	defer func() {
		if err := txHandler.End(err); err != nil {
			err = errors.Wrap(err, "failed to close tx")
		}
	}()

	exist, err := i.service.IsAlreadyExistID(ctx, txHandler, id)
	if err != nil {
		return errors.Wrap(err, "failed to is already exist ID")
	}

	if !exist {
		err = &model.NoSuchDataError{
			PropertyName:    model.IDProperty,
			PropertyValue:   id,
			DomainModelName: model.DomainModelNameUser,
		}
		return errors.Wrap(err, "does not exists ID")
	}

	if err := i.repo.Delete(ctx, txHandler, id); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}
