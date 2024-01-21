Package common

This package contains a struct called `Metric` which has two fields: `Value` and `Timestamp`. The Value field is of type float64 and Timestamp field is of type int64. This struct is used to store data related to metrics.

The function `Copy()` returns a new instance of the Metric struct with the same values as the original object.

### Example Usage:
To use this package, you can create an instance of the Metric struct and set its value. The timestamp is automatically set when the instance is created using the current time. Here's how to do it:

```go
package main
import (
    "time"
)

func main() {
	m := common.Metric{Value: 42}
	newM := m.Copy()
	fmt.Println(newM) // Outputs {42 16007968659} where 16007968659 is the current timestamp in seconds since epoch
}
```

### Documentation:

#### Metric struct
The `Metric` struct has two fields:
- `Value`: type float64, which stores the value of a metric.
- `Timestamp`: type int64, which stores the timestamp in seconds since epoch when the metric was created. It is automatically set to the current time when an instance of Metric is created.

#### Copy() function
The `Copy()` function returns a new instance of the Metric struct with the same values as the original object.

Here's how it works:
1. Take in a pointer to the original `Metric` struct and return a new instance of the struct.
2. Create a new instance of the `Metric` struct, initialize its value and timestamp fields with the same values as the original object.
3. Return the new instance.

Here's an example:
```go
package main
import (
    "time"
)

func main() {
	m := common.Metric{Value: 42}
	newM := m.Copy()
	fmt.Println(newM) // Outputs {42 1607968659} where 1607968659 is the current timestamp in seconds since epoch
}
```
Note that `m` and `newM` are different instances of Metric with the same values. The timestamp field is automatically set when the `Metric` object is created.


#### Benefits:
- This function can be used to create a copy of an existing metric, which can be useful for sending metrics from multiple sources and storing them in a database without losing information about their creation time.
- It allows you to manipulate the values of the metric without affecting other instances of it.

#### Limitations:
- This function only copies the `Value` field, not the entire Metric struct. If more fields are added to the struct in future versions, they will also need to be copied over using this method.
