package wallet

import (
	"strings"
)

type PriceNumber string

func (p *PriceNumber) MarshalJSON() ([]byte, error) {
	text := []byte("\"")

	if p == nil {
		text = append(text, '0')
	} else {
		text = append(text, []byte(*p)...)
	}

	return append(text, '"'), nil
}

func (p *PriceNumber) UnmarshalJSON(b []byte) error {
	*p = PriceNumber(strings.Trim(string(b), "\""))
	return nil
}

type Snapshot struct {
	SnapshotId string      `json:"snapshotId"`
	AssetId    string      `json:"assetId"`
	TraceId    string      `json:"traceId"`
	Amount     PriceNumber `json:"amount"`
	Memo       string      `json:"memo"`
	CreatedAt  int64       `json:"createdAt"`

	CounterUserId string `json:"counterUserId"`

	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	TransactionHash string `json:"transactionHash"`
}
