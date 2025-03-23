package errx

import (
	"context"
	"errors"
)

type CustomError struct {
	base       error
	message    string
	httpCode   int
	customCode int
	ctx        context.Context
}

// New returns a new error with the given message.
func New(msg string) error {
	return &CustomError{
		message: msg,
		ctx:     context.Background(),
	}
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return e.message
}

// Unwrap returns the underlying base error of the CustomError.
// It allows access to the wrapped error for further inspection or handling.
func (e *CustomError) Unwrap() error {
	return e.base
}

// Is reports whether the given error is equal to e.
// It returns true if the two errors are the same, and false otherwise.
func (e *CustomError) Is(err error) bool {
	return errors.Is(e.base, err)
}

// As attempts to set the target to the underlying error if it matches the type.
// It returns true if the target was set, and false otherwise.
func (e *CustomError) As(target interface{}) bool {
	return errors.As(e.Unwrap(), target)
}

// Wrap sets the underlying error of e to the given error and returns the error.
// It is intended to be used to record the cause of the error.
func (e *CustomError) Wrap(err error) error {
	e.base = err
	return err
}

// HTTPCode returns the HTTP status code of the error.
// If the code is not explicitly set through `WithHttpCode`, it returns 0.
func (e *CustomError) HTTPCode() int {
	return e.httpCode
}

// CustomCode returns the custom error code of the error.
// If the code is not explicitly set through `WithCustomCode`, it returns 0.
func (e *CustomError) CustomCode() int {
	return e.customCode
}

// Cause returns the underlying base error of the CustomError.
// It provides access to the original error that caused the CustomError.
func (e *CustomError) Cause() error {
	return e.base
}

// Context returns the context associated with the CustomError.
// If no context was associated with the error, it returns the background context.
func (e *CustomError) Context() context.Context {
	return e.ctx
}

// WithHttpCode sets the HTTP status code of the CustomError to the provided value.
// It returns the updated error with the new HTTP status code, allowing chaining
// of method calls. The HTTP status code can be retrieved later using the HTTPCode
// method.
func (e *CustomError) WithHttpCode(httpCode int) error {
	e.httpCode = httpCode
	return e
}

// WithCustomCode sets the custom error code of the CustomError to the provided value.
// It returns the updated error with the new custom error code, allowing chaining
// of method calls. The custom error code can be retrieved later using the CustomCode
// method.
func (e *CustomError) WithCustomCode(customCode int) error {
	e.customCode = customCode
	return e
}

// WithContext sets the context of the CustomError to the provided context.
// It returns the updated error with the new context, allowing chaining
// of method calls. The context can be retrieved later using the Context
// method.
func (e *CustomError) WithContext(ctx context.Context) error {
	e.ctx = ctx
	return e
}
