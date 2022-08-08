package customerror

type customError struct {
	message string
}

func New(message string) *customError {
	return &customError{
		message: message,
	}
}

func (e *customError) Error() string {
	return e.message
}

var (
	ErrInvalidLongURL   error = &customError{"invalid long url"}
	ErrInvalidLinkID    error = &customError{"invalid link id"}
	ErrInvalidShortPath error = &customError{"invalid short path"}
)
