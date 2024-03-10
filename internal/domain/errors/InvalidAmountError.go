package errors

type InvalidAmountError struct {
	Message string
}

func (e *InvalidAmountError) Error() string {
	return e.Message
}

func NewInvalidAmountError(message string) *InvalidAmountError {
	return &InvalidAmountError{Message: message}
}
