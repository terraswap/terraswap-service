package token

import "github.com/terraswap/terraswap-service/internal/pkg/terraswap"

type SwapRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type SwapResponseDto struct {
	MsgType string `json:"type"`
	Value   struct {
		contract    string
		execute_msg string
		sender      string
	}
}

type TokensResponse = []TokenResponse

type TokenResponse struct {
	terraswap.Token
	Decimals int `json:"decimals"`
}
