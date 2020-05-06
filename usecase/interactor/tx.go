package interactor

import (
	"github.com/ashizaki/go-clean-architecture/domain/model"
	"github.com/pkg/errors"
)

// beginTxErrorMsg generates and returns tx begin error message.
func beginTxErrorMsg(err error) error {
	return errors.WithStack(&model.SQLError{
		BaseErr:       err,
		InvalidReason: "failed to begin tx",
	})
}
