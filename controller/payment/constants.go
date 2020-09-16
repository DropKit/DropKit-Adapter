package payment

type token struct {
	Amount     int64  `json:"amount"`
	PrivateKey string `json:"caller_pk"`
}

type tokenTransfer struct {
	Account    string `json:"user_name"`
	Amount     int64  `json:"amount"`
	PrivateKey string `json:"caller_pk"`
}

type tokenTransferResponse struct {
	Code    int
	Message string
	Hash    string
}

type tokenTransferFailResponse struct {
	Code    int
	Message string
}

type tokenBalance struct {
	PrivateKey string `json:"caller_pk"`
	Account    string `json:"user_name"`
}

type tokenBalanceResponse struct {
	Code    int
	Message string
	Amount  int64
}
