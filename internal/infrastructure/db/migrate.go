package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrateDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

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
	err = m.Up()
	if err != nil {
		return errors.New(fmt.Sprintf("error in up migration: %v", err))
	}

	return nil
}

func NewMigrator(db *GormDatabase) func() error {
	return func() error {
		return Migrate(db)
	}
}
