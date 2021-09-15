package util

type NotFoundError struct {
	message string
}

func NewNotFoundError(mes string) *NotFoundError {
	return &NotFoundError{message: mes}
}

func (e *NotFoundError) Error() string {
	return e.message
}
