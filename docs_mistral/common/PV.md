 # common Package

The `common` package provides basic data structures and functions for the application. This documentation covers the `PV` type and its associated methods.

## PV Type

The `PV` type represents a Power Variable, which holds power-related information such as active power (in kW) and capacity. It is defined as follows:

```go
type PV struct {
	P_kW       float64 `json:"p_kw"`  // Active power in kW
	Capacity_W float64 `json:"capacity_w"`  // Maximum capacity in watts
	Timestamp  int64   `json:"timestamp"`  // Unix timestamp of the data point
}
```

The fields are annotated with JSON tags to facilitate JSON encoding and decoding. The active power (`P_kW`) and maximum capacity (`Capacity_W`) are represented as floating-point numbers, while the timestamp is represented as an Unix integer timestamp.

## Copy Method

The `Copy()` method of a `PV` type returns a new `PV` instance with the same values as the given `PV`. It does not modify the original instance:

```go
func (p *PV) Copy() PV {
	return PV{
		P_kW:       p.P_kW,
		Capacity_W: p.Capacity_W,
	}
}
```

This method is useful when you want to create a new instance of `PV` with the same values as an existing one without modifying the original instance.
