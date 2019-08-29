package api

type ErrorType uint

const (
	ErrorNormal ErrorType = iota
	ErrClient
	ErrImplementation
	ErrRateLimited
	ErrForbidden
)

func (e ErrorType) Error() string {
	return e.String()
}

func (e ErrorType) String() string {
	switch e {
	case ErrorNormal:
		return "command failed"
	case ErrClient:
		return "invalid argument"
	case ErrImplementation:
		return "internal error"
	case ErrRateLimited:
		return "rate limited"
	case ErrForbidden:
		return "request forbidden"
	default:
		return "unknown error code"
	}
}

type Error struct {
	Message string
	Code    ErrorType
}

func (e Error) Error() string {
	return e.Message
}
