package kraken

import "math/big"

// Requests
type getLedgersArgs struct {
	Asset  string `json:"asset"`
	Aclass string `json:"aclass"`
	Type_  string `json:"type"`
	Start  int64  `json:"start"`
	End    int64  `json:"end"`
	Offset int    `json:"ofs"`
}

// Responses
type Response struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

//https://docs.kraken.com/rest/#operation/getAccountBalance
type TickerSymbol string

type AccountBalanceResponse struct {
	Asset map[TickerSymbol]string `json:"balance"`
}

// https://docs.kraken.com/rest/#operation/getLedgers
type LedgerID string

type LedgersResponse struct {
	Ledger map[LedgerID]LedgerInfo `json:"ledger"`
	Count  int                     `json:"count"`
}

type LedgerInfo struct {
	RefID   string    `json:"refid"`
	Time    float64   `json:"time"`
	Type    string    `json:"type"`
	Aclass  string    `json:"aclass"`
	Asset   string    `json:"asset"`
	Amount  big.Float `json:"amount"`
	Fee     big.Float `json:"fee"`
	Balance big.Float `json:"balance"`
}
