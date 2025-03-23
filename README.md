# `errx` - A Custom Error Handling Package for Go

`errx` is a Go package that provides a custom error type (`CustomError`) designed to handle enriched error information
such as HTTP status codes, custom error codes, and context propagation. It supports wrapping errors, accessing the
original cause of the error, and managing contextual information to provide better error handling in Go applications.

## Features

- **Error Wrapping**: Supports wrapping errors with additional context for deeper error inspection.
- **HTTP & Custom Error Codes**: Attach HTTP status codes and custom error codes to errors for better error
  categorization.
- **Context Propagation**: Allows associating a `context.Context` with the error, useful for passing context across API
  calls or long-running processes.
- **Error Inspection**: Implements methods for inspecting the underlying cause of an error and checking if errors match
  certain conditions.

## Installation

To install the `errx` package, run:

```bash
go get github.com/hamidghavidel/errx
```

## Usage

### Creating a New Error

You can create a new error with a simple message. By default, it will use the context.Background() for context:

```go
package main

import (
	"fmt"
	"github.com/hamidghavidel/errx"
)

func main() {
	err := errx.New("Something went wrong")
	fmt.Println(err)
}
```

### Wrapping Errors

You can wrap a base error inside a CustomError to preserve the original error:

```go
baseErr := fmt.Errorf("Original error")
customErr := errx.New("An error occurred")
wrappedErr := customErr.Wrap(baseErr)

fmt.Println(wrappedErr) // Output: An error occurred
```

### Adding HTTP and Custom Error Codes

You can add HTTP status codes and custom error codes to the error:

```go
err := errx.New("Resource not found").
WithHttpCode(404).
WithCustomCode(1001)

fmt.Println("HTTP Code:", err.HTTPCode())  // Output: 404
fmt.Println("Custom Code:", err.CustomCode()) // Output: 1001
```

### Accessing the Context

You can associate a context.Context with the error to propagate request-scoped values:

```go
ctx := context.Background()
err := errx.New("Failed to fetch data").WithContext(ctx)
fmt.Println(err.Context()) // Output: context.Background
```

### Error Inspection

You can inspect errors using Unwrap and Is methods:

```go
baseErr := fmt.Errorf("Original error")
wrappedErr := errx.New("Wrapped error").Wrap(baseErr)

// Unwrap the error
unwrappedErr := wrappedErr.Unwrap()
fmt.Println(unwrappedErr) // Output: Original error

// Use Is to check if the base error matches
fmt.Println(wrappedErr.Is(baseErr)) // Output: true
```

### Method Chaining

The package allows method chaining to set multiple properties of the error:

```go
err := errx.New("Unauthorized").
WithHttpCode(401).
WithCustomCode(2001).
WithContext(ctx)

fmt.Println("Error:", err)
fmt.Println("HTTP Code:", err.HTTPCode())
fmt.Println("Custom Code:", err.CustomCode())
fmt.Println("Context:", err.Context())
```

## Methods Overview

### `New(msg string) error`

Creates a new `CustomError` with the given message. Uses `context.Background()` by default.

### `Error() string`

Returns the error message as a string.

### `Unwrap() error`

Returns the underlying base error, enabling error inspection.

### `Is(target error) bool`

Reports whether the target error matches the current error.

### `As(target interface{}) bool`

Attempts to set the target to the underlying error if it matches the type.

### `Wrap(err error) error`

Wraps the provided error as the underlying error and returns a new `CustomError`.

### `HTTPCode() int`

Returns the HTTP status code associated with the error.

### `CustomCode() int`

Returns the custom error code associated with the error.

### `Cause() error`

Returns the underlying cause of the error.

### `Context() context.Context`

Returns the context associated with the error.

### `WithHttpCode(httpCode int) error`

Sets the HTTP status code for the error and returns the updated error.

### `WithCustomCode(customCode int) error`

Sets the custom error code for the error and returns the updated error.

### `WithContext(ctx context.Context) error`

Sets the context for the error and returns the updated error.

### Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/hamidghavidel/errx"
)

func main() {
	// Create a new error with a message
	err := errx.New("File not found")

	// Wrap another error
	baseErr := fmt.Errorf("Original error")
	wrappedErr := err.Wrap(baseErr)

	// Add HTTP and Custom error codes
	errWithCodes := wrappedErr.
		WithHttpCode(404).
		WithCustomCode(1001)

	// Add context to the error
	ctx := context.Background()
	errWithContext := errWithCodes.WithContext(ctx)

	// Print error details
	fmt.Println("Error:", errWithContext)
	fmt.Println("HTTP Code:", errWithContext.HTTPCode())
	fmt.Println("Custom Code:", errWithContext.CustomCode())
}
```

## License
See the [LICENSE](LICENSE) file for license information.