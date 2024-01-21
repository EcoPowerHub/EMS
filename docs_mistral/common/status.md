 # Package common

This package contains basic data structures and functions used throughout the application.

## Type Definition: Status

```go
type Status struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
```

The `Status` type is a custom data structure, representing an observable state with a corresponding value and timestamp. It is JSON serializable as indicated by the comments.

### Value

Type: `float64`

A numerical value representing the status.

### Timestamp

Type: `int64`

An integer value representing the Unix timestamp when the status was last updated.

## Function Definition: Copy

```go
func (s *Status) Copy() Status {
	return Status{
		Value:     s.Value,
		Timestamp: s.Timestamp,
	}
}
```

This method returns a new `Status` instance with the same values as the receiver but without sharing the same memory. This is an efficient and safe way to create a copy of the `Status` data structure. The copied `Status` struct will have its own unique memory allocation, making it ideal for use cases where multiple instances are being manipulated concurrently or when data immutability is desired.
