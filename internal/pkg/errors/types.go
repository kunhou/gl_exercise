package errors

// BadRequest new BadRequest error
func BadRequest(reason string) *Error {
	return New(400, reason)
}

// NotFound new NotFound error
func NotFound(reason string) *Error {
	return New(404, reason)
}

// InternalServer new InternalServer error
func InternalServer(reason string) *Error {
	return New(500, reason)
}
