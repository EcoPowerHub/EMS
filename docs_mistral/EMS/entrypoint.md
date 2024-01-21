 # ems Package Documentation

The `ems` package is the main entry point of an Energy Management System (EMS) written in Go. It initializes and runs the system according to the provided configuration file.

## Table of Contents
1. [Import Statements](#import-statements)
2. [Variables and Constants](#variables-and-constants)
3. [Core Structure](#core-structure)
4. [Start Function](#start-function)
5. [Import Statements](#import-statements)

## Import Statements <a name="import-statements"></a>

The `ems` package imports the following Go packages:

1. `fmt`: For formatted I/O operations.
2. `log`: For logging errors and messages.
3. `time`: For handling time delays and durations.
4. `github.com/EcoPowerHub/EMS/config`: For reading the configuration file.
5. `github.com/EcoPowerHub/EMS/driver`: For managing equipment drivers.
6. `github.com/EcoPowerHub/context/pkg`: For handling contexts.

## Variables and Constants <a name="variables-and-constants"></a>

There are no constants defined in this package, but there is a global `ems` variable of type `core`.

## Core Structure <a name="core-structure"></a>

The `core` structure contains the configuration, manager, and context instances:

```go
type core struct {
	configuration config.EMS
	manager       *manager.Manager
	context       *context.Context
}
```

## Start Function <a name="start-function"></a>

The `Start` function initializes the EMS by reading the configuration file, creating and setting up the manager and context instances, and starting the main control loop:

```go
func Start(confpath string) {
	// ...
}
```

### Reading Configuration <a name="reading-configuration"></a>

The `Start` function starts by reading the configuration file from the given path:

```go
err = utils.ReadJsonFile(confpath, &ems.configuration)
if err != nil {
	// ...
}
```

### Creating ManagerEquipment and Contexts <a name="creating-managerequipment-and-contexts"></a>

After reading the configuration file, `Start` creates a new managerEquipment instance and context instance:

```go
managerEquipment, err := manager.New(ems.configuration.Equipments)
if err != nil {
	// ...
}

ems.context, err = context.New(ems.configuration.Contexts)
if err != nil {
	// ...
}
```

### Setting Up ManagerEquipment and Contexts <a name="setting-up-managerequipment-and-contexts"></a>

The `Start` function then sets up the managerEquipment and context instances:

```go
ems.manager = managerEquipment
if err := managerEquipment.SetupEquipments(); err != nil {
	// ...
}
```

### Main Control Loop <a name="main-control-loop"></a>

Finally, the `Start` function enters the main control loop that executes all drivers and reads their values every second:

```go
for {
	if err := managerEquipment.InitCycle(); err != nil {
		// ...
	}
	readings := managerEquipment.Read()
	fmt.Printf("Readings %s\n", readings)
	time.Sleep(1 * time.Second)
}
```
