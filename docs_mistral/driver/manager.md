 # driver Package Documentation

The `driver` package is a fundamental part of the EMS (Energy Management System) project developed by EcoPowerHub. It houses the core logic related to managing and interacting with various types of equipment.

## Import Statements

```go
import (
	"fmt"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/config"
	equipment "github.com/EcoPowerHub/EMS/driver/factory"
	"golang.org/x/sync/errgroup"
)
```

The package imports essential packages and custom types that are utilized throughout the `driver` package, such as common constants, configuration settings, equipment factories, and concurrency handling (via `errgroup`).

## Manager Struct

```go
type Manager struct {
	Equipments []equipment.Driver
}
```

The `Manager` type acts as the primary data structure for managing a collection of registered equipment drivers. It maintains an array of all instantiated equipment drivers.

## Initializing a Manager

```go
func New(equipments []config.Equipment) (*Manager, error) {
	var (
		err     error
		manager = &Manager{}
	)
	manager.Equipments, err = equipment.Instanciate(equipments)
	return manager, err
}
```

The `New()` function initializes a new instance of the `Manager` type by instantiating all required equipment drivers using their respective configuration data from the input `equipments` slice. It returns an error if any issue occurs during this process.

## Equipping Functionality

The `driver` package offers several methods for managing and interacting with equipment drivers:

### Reading Equipment Data

```go
func (m *Manager) Read() map[string]map[string]any {
	var read map[string]map[string]any
	for _, d := range m.Equipments {
		read = d.Read()
	}
	return read
}
```

The `Read()` function triggers the `Read()` method on all equipment drivers and collects their returned data into a map of maps. This function can be used to retrieve the current state of all registered equipment in the system.

### Setting Up Equipment

```go
func (m *Manager) SetupEquipments() (err error) {
	for _, e := range m.Equipments {
		if err = e.Configure(); err != nil {
			return
		}
	}
	return
}
```

The `SetupEquipments()` function calls the `Configure()` method on each equipment driver, ensuring all of them are properly configured before further processing.

### Initializing a Cycle

```go
func (m *Manager) InitCycle() (err error) {
	var (
		g = errgroup.Group{}
	)
	for _, e := range m.Equipments {
		if e.State().Value != common.EquipmentStateOnline {
			fmt.Printf("Equipment is not online, skipping")
			continue
		}
		g.Go(e.AddOrRefreshData)
	}
	if err = g.Wait(); err != nil {
		return err
	}
	return
}
```

The `InitCycle()` function is responsible for launching all online equipment drivers' `AddOrRefreshData` method concurrently and then waits for their completion. This method can be used to initiate a new data collection cycle across all registered equipment in the system.
