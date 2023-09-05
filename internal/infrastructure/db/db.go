package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrateDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	err = RunMigrations(database)
	if err != nil {
		log.Fatalf("error in run migrations: %v", err)
	}

	return
}

func RunMigrations(db *GormDatabase) error {
	sqlDB, _ := db.DB.DB()
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
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no change made by migration scripts")
			return nil
		} else {
			return err
		}
	}

	return nil
}
