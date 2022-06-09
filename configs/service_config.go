package configs

import (
	"github.com/spf13/viper"
)

// ServiceConfig contains configs for other services
type ServiceConfig struct {
	TerraSwapApiEndpoint string
}

func serviceConfig(v *viper.Viper) ServiceConfig {
	return ServiceConfig{
		TerraSwapApiEndpoint: v.GetString("services.terraSwapApiEndpoint"),
	}
}
