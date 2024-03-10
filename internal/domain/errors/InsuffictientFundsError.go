package errors


type InsufficientFundsError struct {
	Message string
}

func (e *InsufficientFundsError) Error() string {
	return e.Message
}

func NewInsufficientFundsError(message string) *InsufficientFundsError {
	return &InsufficientFundsError{Message: message}
}
