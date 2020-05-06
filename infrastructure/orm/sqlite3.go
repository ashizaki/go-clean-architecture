package orm

import (
	"fmt"

	"github.com/ashizaki/go-clean-architecture/domain/model"

	"github.com/ashizaki/go-clean-architecture/domain/repository"

	"github.com/ashizaki/go-clean-architecture/infrastructure/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(dbname string) (repository.DbHandler, error) {
	db, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		logger.Logger.Println(fmt.Sprintf("failed to open sqlite3 database. err = %+v", err))
		return nil, err
	}

	migrate(db)

	db = db.Debug()

	return NewDbHandler(db), nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
