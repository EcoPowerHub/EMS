 # common Package Documentation

This package contains types and functions related to Energy Storage System (ESS). The main type defined in this package is `ESS`.

## ESS Type

`type ESS struct {`

- **P_kW** of type `float64` and tagged with json name `"p_kw"`: Represents the power in kilowatt (kW) of the Energy Storage System.
- **Soc** of type `float64` and tagged with json name `"soc"`: Represents the State of Charge (SoC) of the Energy Storage System.
- **Soh** of type `float64` and tagged with json name `"soh"`: Represents the State of Health (SoH) of the Energy Storage System.
- **Capacity_Wh** of type `float64` and tagged with json name `"capacity_wh"`: Represents the total capacity in Wh of the Energy Storage System.
- **Timestamp** of type `int64` and tagged with json name `"timestamp"`: Represents the timestamp in milliseconds.

}`

The `ESS` type also has a method called `Copy()`.

## Copy Method

`func (e *ESS) Copy() ESS {`

This method returns a copy of the received `ESS` instance. It does not modify the original instance and returns a new one with the same values for all fields.

}

With this documentation, users will have a clear understanding of what the package offers, how to use it, and its underlying data structures.
