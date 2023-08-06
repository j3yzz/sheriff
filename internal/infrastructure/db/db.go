package db

import (
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(cfg Config) (database *GormDatabase) {
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
		log.Fatalf("error in connection to mysql: %v", err)
	}

	database = new(GormDatabase)
	database.DB = db

	return
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(NewMigrator),
)
