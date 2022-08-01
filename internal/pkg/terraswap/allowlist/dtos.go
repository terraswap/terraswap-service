package allowlist

type AllowlistToken interface {
	Cw20Allowlist | IbcTokenAllowlist
	Equal(T interface{}) bool
}
type Cw20Allowlist struct {
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Symbol   string `json:"symbol,omitempty" yaml:"symbol,omitempty"`
	Token    string `json:"token,omitempty" yaml:"token,omitempty"`
	Icon     string `json:"icon,omitempty" yaml:"icon,omitempty"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Decimals uint   `json:"decimals,omitempty" yaml:"decimals,omitempty"`
}

func (w Cw20Allowlist) Equal(src interface{}) bool {
	l, ok := src.(Cw20Allowlist)
	return ok && w.Icon == l.Icon && w.Symbol == l.Symbol && w.Token == l.Token &&
		w.Protocol == l.Protocol && w.Decimals == l.Decimals
}

type IbcTokenAllowlist struct {
	// ibc/HASH_FORMAT
	Denom string `json:"denom,omitempty" yaml:"denom"`
	// transfer/channel-x
	Path string `json:"path,omitempty" yaml:"path"`
	// uatom
	BaseDenom string `json:"base_denom,omitempty" yaml:"base_denom"`
	// ATOM
	Symbol string `json:"symbol,omitempty" yaml:"symbol"`
	// Cosmos
	Name     string `json:"name,omitempty" yaml:"name"`
	Icon     string `json:"icon,omitempty" yaml:"icon"`
	Decimals int    `json:"decimals,omitempty" yaml:"decimals"`
}

func (lhs IbcTokenAllowlist) Equal(src interface{}) bool {
	rhs, ok := src.(IbcTokenAllowlist)
	return ok && lhs.Denom == rhs.Denom && lhs.Path == rhs.Path && lhs.BaseDenom == rhs.BaseDenom &&
		lhs.Symbol == rhs.Symbol && lhs.Name == rhs.Name && lhs.Icon == rhs.Icon
}

type Cw20AllowlistResponse struct {
	Mainnet map[string]Cw20Allowlist `json:"mainnet"`
	Classic map[string]Cw20Allowlist `json:"classic"`
	Testnet map[string]Cw20Allowlist `json:"testnet"`
}

type IbcAllowlistResponse struct {
	Mainnet map[string]IbcTokenAllowlist `json:"mainnet"`
	Classic map[string]IbcTokenAllowlist `json:"classic"`
	Testnet map[string]IbcTokenAllowlist `json:"testnet"`
}

type AllowlistResponse interface {
	Cw20AllowlistResponse | IbcAllowlistResponse
}
