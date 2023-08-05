package db

type Config struct {
	MysqlDatabase string `json:"mysqlDatabase" koanf:"MYSQL_DATABASE"`
	MysqlUser     string `json:"mysqlUser" koanf:"MYSQL_USER"`
	MysqlPassword string `json:"mysqlPassword" koanf:"MYSQL_PASSWORD"`
	MysqlPort     string `json:"mysqlPort" koanf:"MYSQL_PORT"`
}
