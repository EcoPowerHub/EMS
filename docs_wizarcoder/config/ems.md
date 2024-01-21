```golang
// Package config defines how is structured the complete configuration file
package config

import (
	"github.com/EcoPowerHub/context/pkg"
)

type EMS struct {
	Equipments []Equipment `json:"equipments"`
	Contexts   pkg.Configuration `json:"context"`
}
```

The package `config` defines a structure that holds the complete configuration file for the application, which is composed of two main components: `Equipments` and `Contexts`. The first one is a slice of `Equipment` struct type, while the second one is an instance of the `pkg.Configuration` struct defined in the `context` package.
