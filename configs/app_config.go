package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// ApiServerConfig is config struct for app
type ApiServerConfig struct {
	Name    string
	Host    string
	Port    int
	Swagger bool
	Mode    string
}

func appConfig(v *viper.Viper) ApiServerConfig {
	return ApiServerConfig{
		Name:    v.GetString("api.name"),
		Host:    v.GetString("api.host"),
		Port:    v.GetInt("api.port"),
		Swagger: v.GetBool("api.swagger"),
		Mode:    v.GetString("api.mode"),
	}
}

func (c *ApiServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
