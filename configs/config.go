package configs

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	fileName  = "config"
	envPrefix = "app"
)

var envConfig Config

// Config aggregation
type Config struct {
	App       ApiServerConfig
	Log       LogConfig
	Service   ServiceConfig
	Sentry    SentryConfig
	Terraswap TerraswapConfig
	Rdb       RdbConfig
	Cache     CacheConfig
}

// Init is explicit initializer for Config
func New() Config {
	v := initViper()
	envConfig = Config{
		App:       appConfig(v),
		Log:       logConfig(v),
		Service:   serviceConfig(v),
		Sentry:    sentryConfig(v),
		Terraswap: terraswapConfig(v),
		Rdb:       rdbConfig(v),
		Cache:     cacheConfig(v),
	}
	return envConfig
}

// Get returns Config object
func Get() Config {
	return envConfig
}

func initViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(fileName)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root := "api/"
	i := strings.LastIndex(path, root)
	if i != -1 {
		path = path[:i+len(root)]
	}
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// All env vars starts with APP_
	v.AutomaticEnv()
	return v
}
