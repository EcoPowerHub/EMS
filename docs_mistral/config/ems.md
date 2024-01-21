 # config Package Documentation

The `config` package is responsible for defining the structure of the complete configuration file in Go. This package imports the `github.com/EcoPowerHub/context/pkg` package for handling contexts.

## Struct Definition

```go
type EMS struct {
	Equipments       []Equipment           `json:"equipments"`
	Contexts         context.Configuration `json:"context"`
}
```

The `EMS` type is the main structure that holds all the configuration data. It consists of two parts:

1. **Equipments**: This is a slice of `Equipment` type that contains information about various equipment in the system.
2. **Contexts**: This field represents the context configurations of the system and is of the `context.Configuration` type.

## Equipment Type

The `Equipment` type, which is not defined within this package, holds the configuration data related to individual pieces of equipment. This may include their names, settings, or other relevant information.

## Contexts Type

The `Contexts` type, imported from the `github.com/EcoPowerHub/context/pkg`, represents the context configurations for the system. The specific fields and types within this struct would depend on how the `context` package is defined.

## Usage

To use the `config` package in your Go project, follow these steps:

1. Import the package at the beginning of your Go file:
   ```go
   import "path/to/config"
   ```
2. Instantiate an instance of the `EMS` type and parse the configuration file into it:
   ```go
   var config EMS
   err := json.Unmarshal([]byte(configData), &config)
   if err != nil {
       log.Fatalf("Error unmarshalling config data: %v", err)
   }
   ```
3. Access the configuration data through the `Equipments` and `Contexts` fields of the `EMS` instance.

For example, to get a list of all equipment names:
```go
names := make([]string, len(config.Equipments))
for i, eqp := range config.Equipments {
    names[i] = eqp.Name
}
```
