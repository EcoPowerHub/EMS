 # Equipment Package

This package provides a set of interfaces and functions to manage equipment drivers in the Energy Management System (EMS). The equipment drivers are responsible for communicating with different types of energy equipment, such as solar panels or wind turbines.

## Imported Packages

- `fmt`: For formatted I/O.
- `github.com/EcoPowerHub/EMS/common`: For the `common.EquipmentState` type.
- `github.com/EcoPowerHub/EMS/config`: For the `config.Equipment` type and configuration handling.
- `github.com/EcoPowerHub/EMS/driver/drivers/generic/PV`: For the `pv.New` function, which creates a new instance of the PV driver.

## Equipment Interface

The `Driver` interface defines methods common to all equipment drivers:

- `Configure() error`: Configures the driver.
- `AddOrRefreshData() error`: Adds or refreshes data from the equipment.
- `State() common.EquipmentState`: Returns the current state of the equipment.
- `Read(map[string]map[string]any)`: Reads data from the equipment.
- `Write(map[string]map[string]any) error`: Writes data to the equipment.

## Instantiating Equipment Drivers

The `Instanciate` function creates and initializes a slice of `Driver` instances based on the given list of configuration objects (`config.Equipment`). It returns an error if any driver instantiation fails or if the list contains unsupported equipment types.

## Creating a New Equipment Driver

The `newDriver` function is used internally by `Instanciate` to create a new equipment driver instance based on the configuration data (`config.Equipment`). It instantiates the driver according to the equipment type ID. Currently, only PV drivers are supported. If an unsupported equipment type ID is encountered, an error is returned.

## Usage Example

To use this package, import it and call `Instanciate` with a list of configuration objects (`config.Equipment`) as follows:

```go
import "github.com/EcoPowerHub/EMS/equipment"

// Configuration slice example
listEquipments := []config.Equipment{
	{Id: "generic/pv", Host: "localhost:8081"},
	{Id: "unknownDriverType", Host: "notReached"},
}

drivers, err := equipment.Instanciate(listEquipments)
if err != nil {
	// Handle the error
}

// Use the drivers as needed
for _, driver := range drivers {
	err := driver.Configure()
	if err != nil {
		// Handle the error
	}
	// Call other methods on the driver instance
}
```
