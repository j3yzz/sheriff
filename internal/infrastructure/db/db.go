package db

import (
	"errors"
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDatabase struct {
	db *gorm.DB
}

func New(cfg Config) (database *GormDatabase, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlUser,
		cfg.MysqlPassword,
		cfg.MysqlHost,
		cfg.MysqlPort,
		cfg.MysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		err = errors.New(fmt.Sprintf("error in connection to mysql: %w", err))
		return
	}

	database = new(GormDatabase)
	database.db = db

	return
}

var Module = fx.Options(fx.Provide(New))
