This package provides an EMS (Electrical Machine Simulator). The `core` struct is the main struct for this package, which contains a configuration and a context. The `Start()` function reads the configuration file and creates a managerEquipment with it. If any error occurs during these processes, it will return the error and stop execution. Then, a while cycle starts to execute all drivers in the managerEquipment and read their values every second, printing them to console.

### Inputs:
- `confpath`: A string representing the path of the configuration file.

### Outputs:
- The package outputs nothing. If an error occurs during its execution, it will be logged and the program will stop.

### Variables:
- `ems core` is the main struct for this package. It contains a configuration and a context.
- `configuration`: A `config.EMS` struct that represents the EMS's configurations.
- `manager`: A pointer to an instance of the manager.Manager type, which manages all equipments in the simulation.
- `context`: A pointer to an instance of the context.Context type, which contains information about all contexts defined in the configuration file.

### Function:
#### Start()
This function reads the given configuration file and creates a managerEquipment with it. It then creates the context object using the `context.New()` function and passes it to the managerEquipment, which initializes all equipments in it. The while cycle starts by calling `managerEquipment.InitCycle()`. If any error occurs during this process, it will return the error and stop execution.

### Code:
```go
package ems

import (
	"fmt"
	"log"
	"time"

	"github.com/EcoPowerHub/EMS/config"
	manager "github.com/EcoPowerHub/EMS/driver"
	"github.com/EcoPowerHub/EMS/utils"
	context "github.com/EcoPowerHub/context/pkg"
)

var ems core

type core struct {
	configuration config.EMS
	manager       *manager.Manager
	context       *context.Context
}

// Start reads the configuration and sets up an EMS simulation.
func Start(confpath string) {
	var (
		err error
	)

	err = utils.ReadJsonFile(confpath, &ems.configuration)
	if err != nil {
		// #8
		log.Fatalf("Error: %s\n", err)
		return
	}

	// Create a managerEquipment with the parsed config
	managerEquipment, err := manager.New(ems.configuration.Equipments)
	if err != nil {
		fmt.Printf("Failed to create managerEquipment: %s\n", err)
		return
	}

	// Create the context
	ems.context, err = context.New(ems.configuration.Contexts)
	if err != nil {
		fmt.Printf("Failed to create contexts: %s\n", err)
		return
	}

	ems.manager = managerEquipment
	if err := managerEquipment.SetupEquipments(); err != nil {
		fmt.Printf("Failed to setup equipments: %s\n", err)
		return
	}

	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := managerEquipment.InitCycle(); err != nil {
			return
		}
		// Reading drivers values
		readings := managerEquipment.Read()
		fmt.Printf("Readings %s\n", readings)
		time.Sleep(1 * time.Second)
	}
}
```
