package config

import (
	"github.com/j3yzz/sheriff/internal/infrastructure/db"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"log"
)

type Config struct {
	fx.Out

	Database db.Config `json:"database" koanf:"database"`
}

func Provide() Config {
	var instance Config

	k := koanf.New("")
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	log.Printf(`
==================================
	Configuration is loaded.
==================================
	`)
	log.Println(k.String("MYSQL_DATABASE"))

	return instance
}
