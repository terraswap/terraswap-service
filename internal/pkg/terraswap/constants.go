package terraswap

const WasmMsgType = "wasm/MsgExecuteContract"

type TokenTypePrefix string

const (
	NativeTokenPrefix TokenTypePrefix = "u"
	IbcTokenPrefix    TokenTypePrefix = "ibc/"
	Cw20TokenPrefix   TokenTypePrefix = "terra1"
)

const (
	ClassicDenomIconUrl = "https://raw.githubusercontent.com/terra-money/assets/master/icon/svg/Terra"
	DenomIconUrl        = "https://raw.githubusercontent.com/terra-money/assets/master/icon/svg"
	IbcIconUrl          = "https://raw.githubusercontent.com/terra-money/assets/master/icon/svg/ibc"
)

const (
	MainnetPrefix = "phoenix"
	ClassicPrefix = "columbus"
	TestnetPrefix = "pisco"
)

const BLOCK_TIME = 6

var contractAddressMap = map[string]ContractsAddress{
	"phoenix-1":  {Factory: "terra1466nf3zuxpya8q9emxukd7vftaf6h4psr0a07srl5zw74zh84yjqxl5qul", Router: "terra13ehuhysn5mqjeaheeuew2gjs785f6k7jm8vfsqg3jhtpkwppcmzqcu7chk"},
	"pisco-1":    {Factory: "terra1jha5avc92uerwp9qzx3flvwnyxs3zax2rrm6jkcedy2qvzwd2k7qk7yxcl", Router: "terra1xp6xe6uwqrspumrkazdg90876ns4h78yw03vfxghhcy03yexcrcsdaqvc8"},
	"columbus-5": {Factory: "terra1ulgw0td86nvs4wtpsc80thv6xelk76ut7a7apj", Router: "terra19f36nz49pt0a4elfkd6x7gxxfmn3apj7emnenf"},
}

type ContractsAddress struct {
	Factory string
	Router  string
}
