package db

type Config struct {
	MysqlHost     string `json:"mysqlHost" koanf:"mysql_host"`
	MysqlDatabase string `json:"mysqlDatabase" koanf:"mysql_database"`
	MysqlUser     string `json:"mysqlUser" koanf:"mysql_user"`
	MysqlPassword string `json:"mysqlPassword" koanf:"mysql_password"`
	MysqlPort     string `json:"mysqlPort" koanf:"mysql_port"`
}
