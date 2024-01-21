Common package documentation

This package provides essential structs and functions to handle the common functionality across all modules in the project.

`common/ess.go` contains an `ESS` struct that represents a Electric Storage System (ESS) data.

```golang
type ESS struct {
	P_kW        float64 `json:"p_kw"`
	Soc         float64 `json:"soc"`
	Soh         float64 `json:"soh"`
	Capacity_Wh float64 `json:"capacity_wh"`
	Timestamp   int64   `json:"timestamp"`
}
```

- `P_kW` represents the power output of the ESS in kW.
- `Soc` represents the state of charge (SOC) of the ESS in percentage.
- `Soh` represents the State of Health of the ESS, which is a measure of how much remaining capacity there is in the battery.
- `Capacity_Wh` represents the total energy storage capacity of the ESS in Watt hours.
- `Timestamp` is the timestamp when the data was collected from the ESS.

The `Copy()` function is used to create a new instance of `ESS` struct with the same values as the original one. This function returns a new instance of `ESS` struct so that any changes made to it will not affect the original instance of `ESS`.

```golang
func (e *ESS) Copy() ESS {
	return ESS{
		P_kW:        e.P_kW,
		Soc:         e.Soc,
		Soh:         e.Soh,
		Capacity_Wh: e.Capacity_Wh,
	}
}
```

This function takes a pointer of `ESS` struct as an input and returns a new instance of the same struct with all its values copied.
