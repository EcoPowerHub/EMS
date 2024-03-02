 # Services Package Documentation

This package contains the `Manager` type responsible for setting up and updating services in the system. It imports necessary packages and defines an interface for a service, a `Manager` struct, and helper functions to create and set up services.

## Imports
```go
import (
	"fmt"
	"sort"

	context "github.com/EcoPowerHub/context/pkg"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/services/peakshaving"
	"github.com/rs/zerolog"
)
```

## Service Interface
The `Service` interface defines the Configure and Update methods for a service.
```go
type Service interface {
	Configure(configuration any, inputsConf any, outputsConf any) error
	Update() error
}
```

## Manager Struct
The `Manager` struct holds the application context, a map of configured services, and a sorted slice of service names based on their priority.
```go
type Manager struct {
	ctx            *context.Context
	conf           map[string]config.Service
	services       map[string]Service
	sortedServices []string
}
```

## New Function
The `New` function initializes a new `Manager` instance with the provided configuration and context.
```go
func New(conf map[string]config.Service, ctx *context.Context) (*Manager, error) {
	return &Manager{
		conf:     conf,
		services: make(map[string]Service),
		ctx:      ctx,
	}, nil
}
```

## SetupServices Function
The `SetupServices` function initializes each service by creating a new instance using the provided id and then configuring it with the corresponding configuration. It also sorts the services based on their priority.
```go
func (m *Manager) SetupServices() error {
	// ...
}
```

## UpdateServices Function
The `UpdateServices` function updates each service in order of priority.
```go
func (m *Manager) UpdateServices() error {
	// ...
}
```

## newService Function
The `newService` function creates a new instance of a specific service based on the provided id.
```go
func (m *Manager) newService(id string) (Service, error) {
	// ...
}
```

## Switch Statement in newService Function
The switch statement initializes a new `peakshaving` service instance when the id is "peakshaving". If the id is not recognized, it returns an error.
```go
func (m *Manager) newService(id string) (Service, error) {
	switch id {
	case "peakshaving":
		// ...
	default:
		// ...
	}
}
```

This document provides an overview of the services package including its imports, interfaces, structures, and functions. The `Manager` struct holds the configuration and manages the creation, configuration, and updating of various services.
