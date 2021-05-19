package response

type Transaction struct {
	Amount            Balance `json:"amount"`
	Fee               Balance `json:"fee"`
	Confirmations     int     `json:"confirmations"`
	Blockhash         string  `json:"blockhash"`
	Blockindex        int64   `json:"blockhash"`
	Blocktime         int64   `json:"blocktime"`
	Txid              int64   `json:"txid"`
	Time              int64   `json:"time"`
	Timereceived      int64   `json:"timereceived"`
	Bip125Replaceable string  `json:"bip125-replaceable"`
}

type TransactionDetails struct {
	Address   string `json:"address"`
	Category  string `json:"category"`
	Vout      int    `json:"vout"`
	Abandoned bool   `json:"abandoned"`
	Label     string `json:"label"`
}
