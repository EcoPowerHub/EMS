 # common package documentation

This package provides utilities and data structures for the application. The most notable type in this package is `Metric`.

## Metric Type

The `Metric` type represents a single metric data point, holding a value and its corresponding timestamp.

```go
type Metric struct {
	Value     float64  `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
```

* `Value`: (float64) represents the numerical value of a metric.
* `Timestamp`: (int64) represents the time in milliseconds when this metric data point was recorded.

## Methods

### Copy method

The `Copy()` method creates and returns a new `Metric` instance with identical values to the existing one. This method is useful for creating copies of metrics without modifying the original ones.

```go
func (m *Metric) Copy() Metric {
	return Metric{
		Value:     m.Value,
		Timestamp: m.Timestamp,
	}
}
```

This method returns a new `Metric` instance with the same values as the receiver without modifying the original metric. This can be useful when making copies of metrics for various purposes within the application, like maintaining immutable data structures.
