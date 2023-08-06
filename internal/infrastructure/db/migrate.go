package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrateDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateError func() error

func Migrate(db *GormDatabase) error {
	sqlDB, _ := db.db.DB()
	instance, err := migrateDriver.WithInstance(sqlDB, &migrateDriver.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("error in connection to mysql: %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		instance,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("error in creating migraton database instance: %v", err))
	}
	_ = m.Up()

	return nil
}

func NewMigrator(db *GormDatabase) MigrateError {
	return func() error {
		return Migrate(db)
	}
}
