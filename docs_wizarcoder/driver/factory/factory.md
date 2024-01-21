The package `equipment` provides functions and structures for managing EMS equipments. It contains the following components:

1. A struct named `Driver` which defines the interface that all drivers of EMS equipments should implement.
2. A function called `Instanciate` which takes a list of equipments as input and returns an array of driver instances.
3. A function called `newDriver` which takes an equipment json object and returns a driver instance based on its ID.

The `Driver` interface has the following methods:
1. Configure method to configure the driver.
2. AddOrRefreshData method to add/refresh data for the driver.
3. State method to get the state of the equipment.
4. Read method to read data from an equipment.
5. Write method to write data to an equipment.

The `EquipmentState` struct is defined in the common package and it contains a field called `Id` which represents the ID of the equipment and other fields are dependent on each driver type.

The package `config` includes the `Equipment` structure that defines the format of an equipment json object. It has two fields:
- Id: string representing the type of driver
- Host: string representing the address of the device.

### Driver interface
```go
type Driver interface {
	Configure() error
	AddOrRefreshData() error
	State() common.EquipmentState
	Read() map[string]map[string]any
	Write(map[string]map[string]any) error
}
```

### newDriver function
The `newDriver` function is used to create a driver instance based on the provided equipment json object. It takes an equipment json object as input and returns a driver instance.
```go
func newDriver(equipmentJson config.Equipment) (Driver, error)
```

### Instanciate function
The `Instanciate` function is used to create driver instances for each equipment in the list of equipments provided as input. It takes an array of `config.Equipment` objects and returns an array of drivers. If there's any error during creation, it returns an error.
```go
func Instanciate(listEquipmentsJson []config.Equipment) ([]Driver, error)
```

The function iterates through the equipments in the input list and creates a driver instance for each of them using `newDriver` function. It returns an array of drivers and any errors that may occur during creation.

### Imported packages
The package uses two imported packages:
- `github.com/EcoPowerHub/EMS/common`: The common package contains the `EquipmentState` struct and other related types.
- `github.com/EcoPowerHub/EMS/config`: The config package defines the equipment json object format.

### Usage example
Here's an example of how to use this package:
```go
package main

import (
	"fmt"
	"github.com/EcoPowerHub/EMS/equipment"
)

func main() {
	// Create a list of equipments
	listEquipments := []config.Equipment{
		{
			Id: "generic/pv",
			Host: "localhost:8080",
		},
		{
			Id: "generic/batery",
			Host: "localhost:9090",
		},
	}

	// Instanciate drivers for the equipments
	drivers, err := equipment.Instanciate(listEquipments)
	if err != nil {
		panic(err)
	}

	// Configure each driver
	for _, driver := range drivers {
		driver.Configure()
	}

	// Get the state of each driver
	for _, driver := range drivers {
		fmt.Println("State:", driver.State())
	}

	// Read data from each driver and print them
	for _, driver := range drivers {
		data, err := driver.Read()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Data:", data)
	}

	// Write some data to each driver
	for _, driver := range drivers {
		writeData := map[string]map[string]any{
			"voltage": {"value": 230},
			"current": {"value": 1.5},
	}
	err = driver.Write(writeData)
	if err != nil {
		panic(err)
	}
}
```
