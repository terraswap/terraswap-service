package terraswap

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type SwapSendMsg struct {
	Swap struct {
		MaxSpread   string `json:"max_spread,omitempty"`
		BeliefPrice string `json:"belief_price,omitempty"`
		Deadline    uint64 `json:"deadline,omitempty"`
	} `json:"swap"`
}

type withdrawLiquidity struct {
	WithdrawLiquidity struct {
		MinAssets []OfferAsset `json:"min_assets,omitempty"`
		Deadline  uint64       `json:"deadline,omitempty"`
	} `json:"withdraw_liquidity"`
}

func GetSwapSendMsg(max_spread, belief_price string, deadline uint64) (SwapSendMsg, error) {
	sendMsg := SwapSendMsg{}
	sendMsg.Swap.MaxSpread = max_spread
	sendMsg.Swap.BeliefPrice = belief_price
	sendMsg.Swap.Deadline = deadline
	return sendMsg, nil
}

func GetWithdrawSendMsg(minAssets []OfferAsset, deadline uint64) (string, error) {
	msg := withdrawLiquidity{}
	msg.WithdrawLiquidity.MinAssets = minAssets
	msg.WithdrawLiquidity.Deadline = deadline
	data, err := json.Marshal(msg)
	if err != nil {
		panic(errors.Wrap(err, "cannot generate swap send message"))
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func AddressToAssetInfo(addr string) (*AssetInfo, error) {
	isNative := IsNativeToken(addr) || IsIbcToken(addr)

	if !isNative && !IsCw20Token(addr) {
		msg := fmt.Sprintf("wrong format address(%s)", addr)
		return nil, errors.New(msg)
	}

	if isNative {
		return &AssetInfo{
			NativeToken: &AssetNativeToken{
				Denom: addr,
			},
		}, nil
	}

	return &AssetInfo{
		Token: &AssetCWToken{
			ContractAddr: addr,
		},
	}, nil
}

func BaseUnsignedTx(addr, sender string) *UnsignedTx {
	return &UnsignedTx{
		Type: WasmMsgType,
		Value: Value{
			Coins:    make([]Coin, 0),
			Contract: addr,
			Sender:   sender,
		},
	}
}

func GetQueryMsg(last Pair) ([]byte, error) {
	type pairsQuery struct {
		Pairs struct {
			Limit      int         `json:"limit,omitempty"`
			StartAfter []AssetInfo `json:"start_after,omitempty"`
		} `json:"pairs,omitempty"`
	}
	queryMsg := pairsQuery{
		struct {
			Limit      int         "json:\"limit,omitempty\""
			StartAfter []AssetInfo "json:\"start_after,omitempty\""
		}{
			Limit:      30,
			StartAfter: last.AssetInfos,
		},
	}

	byteQueryMsg, err := json.Marshal(queryMsg)
	if err != nil {
		return nil, errors.Wrapf(err, "terraswap.util.GetQUeryMsg(%v)", last)
	}
	return byteQueryMsg, nil

}
func IsValidToken(src string) bool {
	return IsNativeToken(src) || IsIbcToken(src) || IsCw20Token(src)
}

func IsNativeToken(src string) bool {
	return strings.HasPrefix(src, string(NativeTokenPrefix)) && src != "umcp"
}

func IsIbcToken(src string) bool {
	return strings.HasPrefix(src, string(IbcTokenPrefix))
}

func IsCw20Token(src string) bool {
	return strings.HasPrefix(src, string(Cw20TokenPrefix))
}

func ToDenomSymbol(denom string) string {
	symbol := ""

	if denom == "uluna" {
		symbol = "Luna"
	} else { // ex) uusd, ukrw ...
		symbol = strings.ToUpper(denom[1:])
	}

	return symbol
}

func GetFactoryAddress(chainId, version string) string {
	key := chainId
	if version != "" {
		key = fmt.Sprintf("%s-%s", chainId, version)
	}
	return contractAddressMap[key].Factory
}

func GetRouterAddress(chainId, version string) string {
	key := chainId
	if version != "" {
		key = fmt.Sprintf("%s-%s", chainId, version)
	}
	return contractAddressMap[key].Router
}

func GetPairKeyByAssets(addr1, addr2 string) string {
	keys := []string{addr1, addr2}
	sort.Strings(keys)

	key := "pair"
	for _, k := range keys {
		key = fmt.Sprintf("%s:%s", k, key)
	}
	return key
}

func IsMainnet(chainId string) bool {
	return strings.HasPrefix(chainId, MainnetPrefix)
}

func IsClassic(chainId string) bool {
	return strings.HasPrefix(chainId, ClassicPrefix)
}

func IsTestnet(chainId string) bool {
	return strings.HasPrefix(chainId, TestnetPrefix)
}
