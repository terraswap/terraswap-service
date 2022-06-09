package configs

import "github.com/spf13/viper"

type TerraswapConfig struct {
	ChainId          string
	GrpcHost         string
	Cw20AllowlistUrl string
	IbcAllowlistUrl  string
	FilterUnverified bool
}

func terraswapConfig(v *viper.Viper) TerraswapConfig {
	config := TerraswapConfig{
		ChainId:          v.GetString("terraswap.chainId"),
		GrpcHost:         v.GetString("terraswap.grpcHost"),
		Cw20AllowlistUrl: v.GetString("terraswap.cw20AllowlistUrl"),
		IbcAllowlistUrl:  v.GetString("terraswap.ibcAllowlistUrl"),
		FilterUnverified: v.GetBool("terraswap.filterUnverified"),
	}

	if config.Cw20AllowlistUrl == "" {
		panic("must provide terraswap.cw20AllowlistUrl")
	}

	if config.ChainId == "" {
		panic("must provice terraswap.chainId")
	}

	if config.GrpcHost == "" {
		panic("must provice terraswap.grpcHost")
	}

	if config.IbcAllowlistUrl == "" {
		panic("must provice terraswap.ibcAllowlistUrl")
	}

	return config
}
