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
	err := errx.New("something went wrong")
	fmt.Println(err)
}
```

### Wrapping Errors

You can wrap a base error inside a CustomError to preserve the original error:

```go
baseErr := fmt.Errorf("original error")
wrappedErr := errx.Wrap(baseErr, "an error occurred")

fmt.Println(wrappedErr) // Output: original error: an error occurred
```

### Adding HTTP and Custom Error Codes

You can add HTTP status codes and custom error codes to the error:

```go
err := errx.New("resource not found", errx.WithHTTPCode(404), errx.WithCustomCode(1001))
var customErr errx.CustomError
if errx.As(err, &customErr) {
    fmt.Println("HTTP Code:", customErr.HTTPCode) // Output: 404
    fmt.Println("Custom Code:", customErr.CustomCode) // Output: 1001
}
```

### Accessing the Context

You can associate a context.Context with the error to propagate request-scoped values:

```go
err := errx.New("resource not found", errx.WithContext(context.Background()))
var customErr errx.CustomError
if errx.As(err, &customErr) {
    fmt.Println("Context:", customErr.CTX) // Output: context.background
}
```


### Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/hamidghavidel/errx"
)

func main() {
  baseErr := fmt.Errorf("original error")
  err := errx.Wrap(
    baseErr,
    "file not found",
    errx.WithHTTPCode(404),
    errx.WithCustomCode(1001),
    errx.WithContext(context.Background()),
  )

  var customErr errx.CustomError
  if errx.As(err, &customErr) {
    fmt.Println("Error:", customErr.Error())
    fmt.Println("HTTP Code:", customErr.HTTPCode)
    fmt.Println("Custom Code:", customErr.CustomCode)
    fmt.Println("Context:", customErr.CTX)
    fmt.Println("base error:", customErr.Cause())
  }
}
```

## License

See the [LICENSE](LICENSE) file for license information.