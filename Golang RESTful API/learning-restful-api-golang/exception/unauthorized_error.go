package exception

type UnauthorizedError struct {
	Error string
}

func NewUnauthorized(error string) UnauthorizedError {
	return UnauthorizedError{Error: error}
}
