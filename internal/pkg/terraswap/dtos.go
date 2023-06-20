package terraswap

import (
	"github.com/pkg/errors"
)

type UnsignedTx struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Value struct {
	Coins      []Coin     `json:"coins"`
	Contract   string     `json:"contract"`
	ExecuteMsg ExecuteMsg `json:"execute_msg"`
	Sender     string     `json:"sender"`
}

type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}

func NewCoin(amount, denom string) Coin {
	return Coin{
		Amount: amount,
		Denom:  denom,
	}
}

type IbcDenomTrace struct {
	Path      string
	BaseDenom string
}

func NewIbcDenomTrace(path, baseDenom string) IbcDenomTrace {
	return IbcDenomTrace{path, baseDenom}
}

type ExecuteMsg struct {
	IncreaseAllowance  *IncreaseAllowanceMsg  `json:"increase_allowance,omitempty"`
	Swap               *SwapMsg               `json:"swap,omitempty"`
	Provide            *ProvideMsg            `json:"provide_liquidity,omitempty"`
	Send               *SendMsg               `json:"send,omitempty"`
	RouteSwapOperation *RouteSwapOperationMsg `json:"execute_swap_operations,omitempty"`
}

type IncreaseAllowanceMsg struct {
	Amount  string `json:"amount"`
	Spender string `json:"spender"` // pair contract address
}

type OfferAsset struct {
	Amount string    `json:"amount"`
	Info   AssetInfo `json:"info"`
}

type RouteSwapOperationMsg struct {
	Operations     []RouteSwapOperation `json:"operations"`
	MinimumReceive string               `json:"minimum_receive"`
	Deadline       uint64               `json:"deadline,omitempty"`
}

type RouteSwapOperation struct {
	NativeSwapOperation *NativeSwapOperation `json:"native_swap,omitempty"`
	TerraSwapOperation  *TerraSwapOperation  `json:"terra_swap,omitempty"`
}
type NativeSwapOperation struct {
	OfferDenom string `json:"offer_denom"`
	AskDenom   string `json:"ask_denom"`
}

type TerraSwapOperation struct {
	OfferAssetInfo AssetInfo `json:"offer_asset_info"`
	AskAssetInfo   AssetInfo `json:"ask_asset_info"`
}

type SwapMsg struct {
	OfferAsset  OfferAsset `json:"offer_asset"`
	BeliefPrice string     `json:"belief_price"`
	MaxSpread   string     `json:"max_spread"`
	Deadline    uint64     `json:"deadline,omitempty"`
}

type ProvideMsg struct {
	Assets            []OfferAsset `json:"assets"`
	SlippageTolerance string       `json:"slippage_tolerance"`
	Deadline          uint64       `json:"deadline,omitempty"`
}

type SendMsg struct {
	Amount   string `json:"amount"`
	Contract string `json:"contract"`

	// NOTE: msg must encoded as Base64 format the predefined structure described bellow
	// struct {
	// 	Swap              *struct{} `json:"swap,omitempty"`
	// 	WithdrawLiquidity *struct{} `json:"withdraw_liquidity,omitempty"`
	// }
	Msg interface{} `json:"msg"`
}

type Tokens struct {
	t []Token
}

func NewTokens(tokenSlice []Token) *Tokens {
	if tokenSlice == nil {
		tokenSlice = []Token{}
	}
	return &Tokens{
		t: tokenSlice,
	}
}

func (ts Tokens) Map() TokensMap {
	res := make(map[string]Token, len(ts.t))
	for _, t := range ts.t {
		res[t.ContractAddr] = t
	}

	return res
}

func (ts Tokens) Slice() []Token {
	return ts.t
}

type TokensMap map[string]Token

func (tm TokensMap) Has(addr string) bool {
	_, ok := tm[addr]
	return ok
}

func (tm TokensMap) Empty() bool {
	return len(tm) == 0
}

func (tm TokensMap) Equal(rhs TokensMap) bool {
	if len(tm) != len(rhs) {
		return false
	}
	for k, lv := range tm {
		rv, ok := rhs[k]
		if !ok {
			return false
		}
		if !lv.Equal(rv) {
			return false
		}
	}
	return true
}

func (tm TokensMap) GetDiffMap(src TokensMap) (addTokensMap, removeTokensMap TokensMap) {
	addTokensMap = make(map[string]Token, 0)
	removeTokensMap = make(map[string]Token, 0)

	// copy to map
	nMap := make(map[string]Token, len(src))

	for k, v := range src {
		nMap[k] = v
	}

	// compare diff
	for k, v := range tm {
		value, exist := nMap[k]
		if !exist {
			removeTokensMap[k] = v
			continue
		}
		if !v.Equal(value) {
			removeTokensMap[k] = v
			addTokensMap[k] = value
			continue
		}
		delete(nMap, k)
	}

	for k, v := range nMap {
		addTokensMap[k] = v
	}

	return addTokensMap, removeTokensMap
}

func (tm TokensMap) Append(nMap TokensMap) TokensMap {
	for k, v := range nMap {
		tm[k] = v
	}

	return tm
}

func (t TokensMap) ToTokens() *Tokens {
	res := []Token{}
	for _, v := range t {
		res = append(res, v)
	}

	return &Tokens{
		t: res,
	}
}

type Token struct {
	Name         string `json:"name,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Decimals     int    `json:"decimals,omitempty"`
	TotalSupply  string `json:"total_supply,omitempty"`
	ContractAddr string `json:"contract_addr,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Verified     bool   `json:"verified"`
	Protocol     string `json:"protocol,omitempty"`
}

func (lhs *Token) Equal(rhs Token) bool {
	return lhs.Name == rhs.Name && lhs.Symbol == rhs.Symbol && lhs.Decimals == rhs.Decimals &&
		lhs.TotalSupply == rhs.TotalSupply && lhs.ContractAddr == rhs.ContractAddr &&
		lhs.Icon == rhs.Icon && lhs.Verified == rhs.Verified && lhs.Protocol == rhs.Protocol
}

type TokenType string

const (
	NativeTokenType TokenType = TokenType("NativeToken")
	Cw20TokenType   TokenType = TokenType("Cw20Token")
)

type AssetCWToken struct {
	ContractAddr string `json:"contract_addr,omitempty"`
}
type AssetNativeToken struct {
	Denom string `json:"denom,omitempty"`
}

type AssetInfo struct {
	Token       *AssetCWToken     `json:"token,omitempty"`
	NativeToken *AssetNativeToken `json:"native_token,omitempty"`
}

func (a *AssetInfo) GetKey() string {
	if a.NativeToken != nil {
		return a.NativeToken.Denom
	}
	return a.Token.ContractAddr
}

func (a *AssetInfo) GetTokenType() TokenType {
	if a.NativeToken != nil {
		return NativeTokenType
	}
	return Cw20TokenType
}

type Pairs struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	AssetInfos     []AssetInfo `json:"asset_infos,omitempty"`
	LiquidityToken string      `json:"liquidity_token,omitempty"`
	ContractAddr   string      `json:"contract_addr,omitempty"`
}

func (p *Pair) GetPairKey() string {
	return p.ContractAddr
}

func (p *Pair) GetPairCompositeKey() string {
	if len(p.AssetInfos) != 2 {
		panic(errors.New("pair must have two assets"))
	}
	return GetPairKeyByAssets(p.AssetInfos[0].GetKey(), p.AssetInfos[1].GetKey())
}

type PoolInfo struct {
	Assets []struct {
		Info   AssetInfo
		Amount string `json:"amount"`
	} `json:"assets"`
	TotalShare string `json:"total_share"`
}
