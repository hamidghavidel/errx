package errx

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type Property func(err error) error

type CustomError struct {
	base       error
	Message    string
	HTTPCode   int
	CustomCode int
	CTX        context.Context
}

// Error returns a formatted string representation of the CustomError.
// It concatenates the error message from the base error (if available)
// with the message of the CustomError itself, separated by a colon.
func (e CustomError) Error() string {
	baseErrMsg := ""
	if e.base != nil {
		baseErrMsg = e.base.Error()
	}
	return fmt.Sprintf("%s: %s", baseErrMsg, e.Message)
}

// Cause returns the underlying base error of the CustomError.
// It provides access to the original error that caused the CustomError.
func (e CustomError) Cause() error {
	return e.base
}

// New creates a new error with the given message and applies the given properties.
// If no properties are given, it will simply return a wrapped error with the given message.
// Otherwise, it will apply the properties to the error and return the modified error.
func New(msg string, properties ...Property) error {
	if len(properties) == 0 {
		return errors.New(msg)
	}

	var result error = CustomError{
		Message: msg,
		CTX:     context.Background(),
	}

	for _, property := range properties {
		result = property(result)
	}

	return result
}

// Wrap wraps the given error with the given message and applies the given properties.
// If the given error is a CustomError, it wraps the error and applies the properties.
// Otherwise, it uses errors.Wrap to wrap the error with the given message.
func Wrap(err error, msg string, properties ...Property) error {
	if err == nil {
		return nil
	}

	var customErr CustomError
	if errors.As(err, &customErr) {
		var result error = &CustomError{
			base:    customErr,
			Message: msg,
		}

		for _, property := range properties {
			result = property(result)
		}

		return result
	}

	return errors.Wrap(err, msg)
}

// WithHTTPCode returns a Property that sets the HTTP code of an error.
// If the error is a CustomError, it updates the HTTPCode of the existing error.
// Otherwise, it creates a new CustomError with the specified HTTP code.
func WithHTTPCode(httpCode int) Property {
	return func(err error) error {
		var customErr CustomError
		if errors.As(err, &customErr) {
			customErr.HTTPCode = httpCode

			return customErr
		}

		return CustomError{
			Message:  err.Error(),
			HTTPCode: httpCode,
		}
	}
}

// WithCustomCode returns a Property that sets the custom code of an error.
// If the error is a CustomError, it updates the CustomCode of the existing error.
// Otherwise, it creates a new CustomError with the specified custom code.
func WithCustomCode(customCode int) Property {
	return func(err error) error {
		var customErr CustomError
		if errors.As(err, &customErr) {
			customErr.CustomCode = customCode

			return customErr
		}

		return CustomError{
			Message:    err.Error(),
			CustomCode: customCode,
		}
	}
}

// WithContext returns a Property that sets the context of an error.
// If the error is a CustomError, it updates the CTX of the existing error.
// Otherwise, it creates a new CustomError with the specified context.
func WithContext(ctx context.Context) Property {
	return func(err error) error {
		var customErr CustomError
		if errors.As(err, &customErr) {
			customErr.CTX = ctx

			return customErr
		}

		return CustomError{
			Message: err.Error(),
			CTX:     ctx,
		}
	}
}

// Is reports whether any error in err's chain has the same value as target.
// This function is a wrapper around pkg/errors.Is.
func Is(err, target error) bool { return errors.Is(err, target) }

// As reports whether the error err satisfies the predicate target.
// This function is a wrapper around pkg/errors.As.
func As(err error, target any) bool { return errors.As(err, target) }
