package wallet

type Snapshot struct {
	SnapshotId string  `json:"snapshotId"`
	AssetId    string  `json:"assetId"`
	TraceId    string  `json:"traceId"`
	Amount     float64 `json:"amount"`
	Memo       string  `json:"memo"`
	CreatedAt  int64   `json:"createdAt"`

	CounterUserId string `json:"counterUserId"`

	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	TransactionHash string `json:"transactionHash"`
}
