package controller

import (
	"github.com/ashizaki/go-clean-architecture/domain/model"
	"github.com/ashizaki/go-clean-architecture/infrastructure/logger"
	"gopkg.in/go-playground/validator.v8"
)

// handleValidatorErr handle validator error.
func handleValidatorErr(err error) error {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		logger.Logger.Println("failed to assert ValidationErrors")
	}

	errs := &model.InvalidParamsError{}

	for _, v := range errors {
		e := &model.InvalidParamError{
			BaseErr:       err,
			PropertyName:  model.PropertyName(v.Field),
			PropertyValue: v.Value,
		}

		errs.Errors = append(errs.Errors, e)
	}

	if len(errs.Errors) == 1 {
		return errs.Errors[0]
	}

	return errs
}
