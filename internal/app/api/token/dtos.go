package token

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
