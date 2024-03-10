package model

type DepositRequestDTO struct {
	Amount string `json:"amount"`
}

type WithdrawRequestDTO struct {
	Amount string `json:"amount"`
}

type NewWalletResponseDTO struct {
	ID string `json:"id"`
	Balance string `json:"balance"`
}

type ErrorResponseDTO struct {
	Message string `json:"message"`
}
