package syserror

type Error interface {
	Code() int
	Error() string
	ErrorReson() error
}

func New(msg string, reason error) Error {
	return &UnknowError{msg, reason}
}
