package orm

import (
	"github.com/ashizaki/go-clean-architecture/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type handler struct {
	db *gorm.DB
}

func (h *handler) Error() error {
	return h.db.Error
}

func (h *handler) Create(value interface{}) repository.SqlHandler {
	return &handler{db: h.db.Create(value)}
}

func (h *handler) Select(out interface{}) repository.SqlHandler {
	return &handler{db: h.db.Find(out)}
}

func (h *handler) Where(query interface{}, args ...interface{}) repository.SqlHandler {
	return &handler{db: h.db.Where(query, args)}
}

func (h *handler) Get(out interface{}, where ...interface{}) repository.SqlHandler {
	return &handler{db: h.db.First(out, where)}
}

func (h *handler) Save(value interface{}) repository.SqlHandler {
	return &handler{db: h.db.Save(value)}
}

func (h *handler) Delete(value interface{}, where ...interface{}) repository.SqlHandler {
	return &handler{db: h.db.Delete(value, where)}
}

func (h *handler) Exec(sql string, values ...interface{}) repository.SqlHandler {
	return &handler{db: h.db.Exec(sql, values)}
}

func (h *handler) Limit(limit interface{}) repository.SqlHandler {
	return &handler{db: h.db.Limit(limit)}
}

func (h *handler) Offset(offset interface{}) repository.SqlHandler {
	return &handler{db: h.db.Offset(offset)}
}

func (h *handler) IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func (h *handler) End(err error) error {
	if p := recover(); p != nil { // rewrite panic
		err = h.Rollback()
		err = errors.Wrap(err, "failed to roll back")
		panic(p)
	} else if err != nil {
		err = h.Rollback()
		err = errors.Wrap(err, "failed to roll back")
	} else {
		err = h.Commit()
		err = errors.Wrap(err, "failed to commit")
	}
	return err
}

func (h *handler) Begin() (repository.TxHandler, error) {
	tx := h.db.Begin()
	return &handler{db: tx}, tx.Error
}

func (h *handler) Commit() error {
	return h.db.Commit().Error
}

func (h *handler) Rollback() error {
	return h.db.Rollback().Error
}

func NewDbHandler(db *gorm.DB) repository.DbHandler {
	return &handler{db: db}
}
