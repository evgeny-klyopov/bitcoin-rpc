package response

type TransactionListSinceBlock struct {
	Transaction
	TransactionDetails
	Account string `json:"account"`
	Comment string `json:"comment"`
	To      string `json:"to"`
}

type ListSinceBlock struct {
	Transactions []TransactionListSinceBlock `json:"transactions"`
	Removed      []TransactionListSinceBlock `json:"removed"`
	Lastblock    string                      `json:"lastblock"`
}
