 # Package pv

This package contains a `Equipment` type that represents a Power Variable (PV) equipment, which is used to read and store power data from a Modbus device. It interfaces with the common `EquipmentState` and `io.PV` types.

## Imports
```go
import (
	"fmt"
	"time"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/common/io"
	"github.com/rs/zerolog"
	"github.com/simonvetter/modbus"
)
```
## Type: Equipment

The `Equipment` type holds the Modbus client, current state, host information, readings and last reading time.

```go
type Equipment struct {
	logger   zerolog.Logger
	mc       *modbus.ModbusClient
	state    common.EquipmentState
	host     string
	readings readings
	lastRead time.Time
}
```
### Functions

#### `AddOrRefreshData() error`

This function attempts to read the power data (register 0) from the Modbus device and updates the equipment state accordingly. It also stores the current reading in the `readings` field and updates the `lastRead` time.

```go
func (e *Equipment) AddOrRefreshData() error {
	// ...
}
```
#### `New(host string) *Equipment`

This constructor creates a new `Equipment` instance with the given host. The initial state is set to `common.EquipmentStateInit`.

```go
func New(host string) *Equipment {
	// ...
}
```
#### `Configure() error`

This function sets up the Modbus client, opens the connection and sets the equipment state accordingly.

```go
func (e *Equipment) Configure() error {
	// ...
}
```
#### `State() common.EquipmentState`

This function returns the current state of the Equipment.

```go
func (e *Equipment) State() common.EquipmentState {
	return e.state
}
```
#### `Read() map[string]map[string]any`

This function returns the current readings as a `map[string]map[string]any`, where the outer key is the io.PV and the inner key is the specific reading, in this case, the power (kW).

```go
func (e *Equipment) Read() map[string]map[string]any {
	// ...
}
```
#### `Write(_ map[string]map[string]any) error`

This function is currently empty, as this Equipment type does not support writing.

```go
func (e *Equipment) Write(_ map[string]map[string]any) error {
	return nil
}
```
## Constants and Variables

### `readings`

A nested type that holds the power reading (float64) for this Equipment.

```go
type readings struct {
	p_w float64
}
```
