package response

type GetTransactionDetails struct {
	TransactionDetails
	Amount Balance `json:"amount"`
}

type GetTransaction struct {
	Transaction
	Hex     string                  `json:"hex"`
	Details []GetTransactionDetails `json:"details"`
}
