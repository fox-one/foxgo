package wallet

const (
	USDT = "815b0b1a-2764-3736-8faa-42d694fa620a" // "USDT",
	BTC  = "c6d0c728-2624-429b-8e0d-d9d19b6592fa" // "BTC",
	BCH  = "fd11b6e3-0b87-41f1-a41f-f0e9b49e5bf0" // "BCH",
	EOS  = "6cfe566e-4aad-470b-8c9a-2fd35b49c68d" //: "EOS",
	ETH  = "43d61dcd-e413-450d-80b8-101d5e903357" //: "ETH",
	ETC  = "2204c1ee-0ea2-4add-bb9a-b3719cfff93a"
	LTC  = "76c802a2-7c88-447f-a93e-c29c9e5dd9c8"
)

type Asset struct {
	AssetId   string  `json:"assetId"`
	AssetKey  string  `json:"assetKey,omitempty"`
	ChainId   string  `json:"chainId"`
	Icon      string  `json:"icon"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	Balance   float64 `json:"balance"`
	PublicKey string  `json:"publicKey,omitempty"`
}

type Assets []Asset
