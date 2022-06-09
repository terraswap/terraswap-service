package configs

import "github.com/spf13/viper"

// db contains configs for other services
type RdbConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func rdbConfig(v *viper.Viper) RdbConfig {
	return RdbConfig{
		Host:     v.GetString("rdb.host"),
		Port:     v.GetString("rdb.port"),
		Database: v.GetString("rdb.database"),
		Username: v.GetString("rdb.username"),
		Password: v.GetString("rdb.password"),
	}
}
