package configs

import "github.com/spf13/viper"

type CacheConfig struct {
	// x.y.z:pppp, no http://
	Host string
	// Database index
	Database int
	// whether tls is required. (true for dev, staging, prod)
	TLSRequired bool
	// CA cert for TLS connection
	CACert string `json:"-"`
	// whether authentication is required. (true for staging, prod)
	AuthRequired bool
	// password for authentication
	Password string `json:"-"`
}

func cacheConfig(v *viper.Viper) CacheConfig {
	return CacheConfig{
		Host:         v.GetString("terraswap.cache.host"),
		Database:     v.GetInt("terraswap.cache.db"),
		TLSRequired:  v.GetBool("terraswap.cache.tlsRequired"),
		CACert:       v.GetString("terraswap.cache.ca"),
		AuthRequired: v.GetBool("terraswap.cache.authRequired"),
		Password:     v.GetString("terraswap.cache.password"),
	}
}
