package controller

import (
	"net/http"
	"strconv"

	"github.com/ashizaki/go-clean-architecture/domain/model"
	"github.com/ashizaki/go-clean-architecture/infrastructure/logger"
	"github.com/ashizaki/go-clean-architecture/usecase/interactor"
	"github.com/pkg/errors"
)

type UserController interface {
	CreateUser(c Context)
	ListUsers(c Context)
	GetUser(c Context)
	UpdateUser(c Context)
	DeleteUser(c Context)
}

type userController struct {
	interactor interactor.UserInteractor
}

func NewUserController(interactor interactor.UserInteractor) UserController {
	return &userController{interactor: interactor}
}

func (c *userController) CreateUser(ctx Context) {
	u := &model.User{}
	if err := ctx.Bind(u); err != nil {
		err = handleValidatorErr(err)
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to bind json"))
		return
	}

	user, err := c.interactor.Store(ctx.Request().Context(), u)
	if err != nil {
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to store user"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) ListUsers(ctx Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = defaultLimit
	}

	cursorInt, err := strconv.Atoi(ctx.Query("cursor"))
	if err != nil {
		cursorInt = defaultCursor
	}

	cursor := uint(cursorInt)

	list, err := c.interactor.Fetch(ctx.Request().Context(), limit, cursor)
	if err != nil {
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to fetch users"))
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *userController) GetUser(ctx Context) {
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || idInt < 1 {
		logger.Logger.Println(err.Error())
		err = &model.InvalidParamError{
			BaseErr:       err,
			PropertyName:  "id",
			InvalidReason: "id should be number and over 0",
		}
		ResponseAndLogError(ctx, err)
		return
	}

	id := uint(idInt)

	user, err := c.interactor.Get(ctx.Request().Context(), id)
	if err != nil {
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to get user"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) UpdateUser(ctx Context) {
	u := &model.User{}
	if err := ctx.Bind(u); err != nil {
		err = handleValidatorErr(err)
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to bind json"))
		return
	}

	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = &model.InvalidParamError{
			BaseErr:       err,
			PropertyName:  "id",
			PropertyValue: ctx.Param("id"),
		}
		err = handleValidatorErr(err)
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to change id from string to int"))
		return
	}

	id := uint(idInt)

	user, err := c.interactor.Update(ctx.Request().Context(), id, u)
	if err != nil {
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to update user"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) DeleteUser(ctx Context) {
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = &model.InvalidParamError{
			BaseErr:       err,
			PropertyName:  "id",
			PropertyValue: ctx.Param("id"),
		}
		err = handleValidatorErr(err)
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to change id from string to int"))
		return
	}

	id := uint(idInt)

	err = c.interactor.Delete(ctx.Request().Context(), id)
	if err != nil {
		ResponseAndLogError(ctx, errors.Wrap(err, "failed to delete comment"))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
