# Slog - Go's built-in Structured Logger

Structured logging has become a cornerstone of robust and maintainable applications. With the introduction of the `slog` package in Go 1.21, developers now have access to a powerful and flexible logging solution that addresses the limitations of the traditional `log` package. This guide talks briefly about `slog`, providing insights and examples.

## What is slog?

The `slog` package is part of Go's standard library. The `slog` package offer structured logging capabilities. Unlike the older `log` package, which outputs plain text logs, `slog` enables developers to log structured data in a way that is easily parsed by downstream systems, such as log aggregation tools and monitoring platforms.

## Key Features

- **Structured Logging**: Log entries can include key-value pairs, making it easier to filter and analyze logs.
- **Multiple Log Levels**: Supports `debug`, `info`, `warn`, `error`, and `fatal` levels for granular control.
- **Flexible Formatting**: Logs can be formatted in JSON, text, or custom formats.
- **Context Propagation**: Allows for context to be passed through log entries, which is crucial for distributed tracing.

## Example Usage

Here's a simple example demonstrating how to use `slog` in a Go application:

```go
package main

import (
    "log/slog"
    "os"
)

func main() {
    // Create a logger that writes to stdout
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // Log an info message with a field
    logger.Info("Application started", "env", "production")

    // Log an error message
    logger.Error("Failed to connect", "error", "connection timeout")
}
```

This example initializes a logger that writes logs to the standard output. It then logs an `info` message with a field named `env` and an `error` message with a field named `error`.
This uses the logfmt logging format. See this example on what the info log above outputs:

```
time=2022-11-02T21:55:48.012+00:00 level=INFO msg="Application started" env="production"
```

### Setting the default logger

Sometimes in projects people tend to use the default logger from a package. This is also possible with `slog`:

```go
package main

import (
    "log/slog"
    "os"
)
func main() {
    slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}
```

now in other files the default logger can simply be called like this:

```go
package other

import "log/slog"

func foo() {
    slog.Info("Hello world")
}
```

## Conclusion

The `slog` package is a step forward in Go's logging ecosystem, offering developers a structured, flexible and powerful tool for logging built-in to the standard library.
