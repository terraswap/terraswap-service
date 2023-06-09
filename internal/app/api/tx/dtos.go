package tx

type SwapTxRequest struct {
	From        string `form:"from" binding:"required"`
	To          string `form:"to" binding:"required"`
	Amount      string `form:"amount" binding:"required"`
	Sender      string `form:"sender"`
	MaxSpread   string `form:"max_spread"`
	BeliefPrice string `form:"belief_price"`
	Deadline    uint64 `form:"deadline"`
	HopCount    int    `form:"hop_count, default=2"`
}

type ProvideTxRequest struct {
	From       string `form:"from" binding:"required"`
	To         string `form:"to" binding:"required"`
	FromAmount string `form:"fromAmount" binding:"required"`
	ToAmount   string `form:"toAmount" binding:"required"`
	Slippage   string `form:"slippage"`
	Sender     string `form:"sender"`
	Deadline   uint64 `form:"deadline"`
}

type WithdrawTxRequest struct {
	LpAddr   string `form:"lpAddr" binding:"required"`
	Amount   string `form:"amount" binding:"required"`
	Sender   string `form:"sender"`
	Deadline uint64 `form:"deadline"`
}

type SwapTxResponse struct {
	MsgType string `json:"type"`
	Value   struct {
		contract    string
		execute_msg string
		sender      string
	}
}
