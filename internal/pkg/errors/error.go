package errors

// Error error
type Error struct {
	Code   int
	Reason string
	Err    error
	Stack  string
}

// New new error
func New(code int, reason string) *Error {
	return &Error{Code: code, Reason: reason}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Reason
}

// WithStack with stack
func (e *Error) WithStack() *Error {
	e.Stack = LogStack(2, 0)
	return e
}
