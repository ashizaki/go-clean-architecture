package repository

type DbHandler interface {
	SqlHandler
	Beginner
}

type SqlHandler interface {
	Error() error
	Create(value interface{}) SqlHandler
	Select(out interface{}) SqlHandler
	Where(query interface{}, args ...interface{}) SqlHandler
	Get(out interface{}, where ...interface{}) SqlHandler
	Save(value interface{}) SqlHandler
	Delete(value interface{}, where ...interface{}) SqlHandler
	Exec(sql string, values ...interface{}) SqlHandler
	Limit(limit interface{}) SqlHandler
	Offset(offset interface{}) SqlHandler
	IsRecordNotFoundError(error) bool
}

type TxHandler interface {
	SqlHandler
	End(error) error
	Commit() error
	Rollback() error
}

type Beginner interface {
	Begin() (TxHandler, error)
}
