package configs

import "github.com/spf13/viper"

type TerraswapConfig struct {
	ChainId          string
	Version          string
	GrpcHost         string
	Cw20AllowlistUrl string
	IbcAllowlistUrl  string
	FilterUnverified bool
}

func terraswapConfig(v *viper.Viper) TerraswapConfig {
	config := TerraswapConfig{
		ChainId:          v.GetString("terraswap.chainId"),
		Version:          v.GetString("terraswap.version"),
		GrpcHost:         v.GetString("terraswap.grpcHost"),
		Cw20AllowlistUrl: v.GetString("terraswap.cw20AllowlistUrl"),
		IbcAllowlistUrl:  v.GetString("terraswap.ibcAllowlistUrl"),
		FilterUnverified: v.GetBool("terraswap.filterUnverified"),
	}

	if config.Cw20AllowlistUrl == "" {
		panic("must provide terraswap.cw20AllowlistUrl")
	}

	if config.ChainId == "" {
		panic("must provide terraswap.chainId")
	}

	if config.GrpcHost == "" {
		panic("must provide terraswap.grpcHost")
	}

	if config.IbcAllowlistUrl == "" {
		panic("must provide terraswap.ibcAllowlistUrl")
	}

	return config
}
