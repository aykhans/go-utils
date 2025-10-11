# go-utils

A collection of generic utility functions for Go projects.

## Installation

```bash
go get github.com/aykhans/go-utils
```

## Packages

### common

Generic type conversion utilities.

**ToPtr** - Convert any value to a pointer
```go
import "github.com/aykhans/go-utils/common"

num := 42
ptr := common.ToPtr(num)  // *int
```

### parser

String parsing utilities with generic type support.

**ParseString** - Parse string to various types
```go
import "github.com/aykhans/go-utils/parser"

num, err := parser.ParseString[int]("42")
duration, err := parser.ParseString[time.Duration]("5s")
isValid, err := parser.ParseString[bool]("true")
```

**ParseStringOrZero** - Parse string or return zero value on error
```go
num := parser.ParseStringOrZero[int]("invalid")  // returns 0
```

**ParseStringWithDefault** - Parse string with default value fallback
```go
num, err := parser.ParseStringWithDefault("invalid", 10)  // returns 10, error
```

**ParseStringOrDefault** - Parse string or return default value without error
```go
num := parser.ParseStringOrDefault("invalid", 10)  // returns 10
```

Supported types: `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float64`, `bool`, `time.Duration`, `url.URL`

### slice

Slice manipulation utilities.

**Cycle** - Create an infinite cycler through items
```go
import "github.com/aykhans/go-utils/slice"

next := slice.Cycle(1, 2, 3)
fmt.Println(next())  // 1
fmt.Println(next())  // 2
fmt.Println(next())  // 3
fmt.Println(next())  // 1
```

**RandomCycle** - Cycle through items with randomization
```go
next := slice.RandomCycle(nil, "a", "b", "c")
// Cycles through all items, then starts from a random position
```

### maps

Map utility functions.

**InitMap** - Initialize a map pointer if nil
```go
import "github.com/aykhans/go-utils/maps"

var m map[string]int
maps.InitMap(&m)
m["key"] = 42  // safe to use
```

**UpdateMap** - Merge entries from one map into another
```go
old := map[string]int{"a": 1, "b": 2}
new := map[string]int{"b": 3, "c": 4}
maps.UpdateMap(&old, new)
// old is now: {"a": 1, "b": 3, "c": 4}
```

### errors

Advanced error handling utilities.

**HandleError** - Process errors with custom matchers
```go
import "github.com/aykhans/go-utils/errors"

handled, result := errors.HandleError(err,
    errors.OnSentinelError(io.EOF, func(e error) error {
        return nil  // EOF is expected, ignore it
    }),
    errors.OnCustomError(func(e *CustomError) error {
        return fmt.Errorf("custom error: %w", e)
    }),
)
```

**HandleErrorOrDie** - Handle errors or panic if unhandled
```go
result := errors.HandleErrorOrDie(err,
    errors.OnSentinelError(context.Canceled, func(e error) error {
        return fmt.Errorf("operation canceled")
    }),
)  // Panics if err doesn't match any handler
```

**HandleErrorOrDefault** - Handle errors with a default fallback
```go
result := errors.HandleErrorOrDefault(err,
    func(e error) error {
        // Default handler for unmatched errors
        return fmt.Errorf("unexpected error: %w", e)
    },
    errors.OnSentinelError(context.Canceled, func(e error) error {
        return fmt.Errorf("operation canceled")
    }),
    errors.OnCustomError(func(e *ValidationError) error {
        return fmt.Errorf("validation failed: %w", e)
    }),
)

// Pass nil to suppress unmatched errors
result := errors.HandleErrorOrDefault(err, nil,
    errors.OnSentinelError(io.EOF, func(e error) error {
        return errors.New("EOF handled")
    }),
)  // Returns nil for unmatched errors
```

**OnSentinelError** - Create matcher for sentinel errors (like `io.EOF`)
```go
matcher := errors.OnSentinelError(io.EOF, func(e error) error {
    log.Println("reached end of file")
    return nil
})
```

**OnCustomError** - Create matcher for custom error types
```go
type ValidationError struct {
    Field string
    Msg   string
}

matcher := errors.OnCustomError(func(e *ValidationError) error {
    log.Printf("validation failed on field %s", e.Field)
    return fmt.Errorf("invalid input: %w", e)
})
```

## Requirements

- Go 1.25.0 or higher
