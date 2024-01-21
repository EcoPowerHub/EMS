 # Package common

This package contains constant definitions and type for equipment state in our Golang application.

## Constants

The following are the defined constants representing different equipment states:

### EquipmentStateInit

```go
const EquipmentStateInit = 1
```

Initial state of the equipment, indicating that it is yet to be initialized or prepared for usage.

### EquipmentStateOnline

```go
const EquipmentStateOnline = 2
```

Equipment is functioning correctly and providing its intended services to users or the system.

### EquipmentStateUnreachable

```go
const EquipmentStateUnreachable = 3
```

Equipment is currently not reachable due to various reasons such as network connectivity issues, power outages, or hardware failures.

### EquipmentStateOffline

```go
const EquipmentStateOffline = 4
```

Equipment is intentionally taken offline for maintenance, upgrades, or repairs.

### EquipmentStateError

```go
const EquipmentStateError = 5
```

Equipment state indicating an error condition. This can occur due to hardware failures, software errors, or other unforeseen circumstances.

## Type: EquipmentState

```go
type EquipmentState struct {
	Value uint8 `json:"value"`
}
```

This type defines an EquipmentState structure which holds the value of a specific equipment state as an enumerated constant (uint8). It is also JSON serialized and deserialized using the "value" field.
