 # utils Package

The `utils` package is a collection of utility functions for handling common tasks such as reading JSON files.

## Overview

The `ReadJsonFile` function in this package is responsible for reading a JSON file and unmarshalling the content into a Go interface. It takes two arguments:

1. `path string` - The path to the JSON file to be read.
2. `target interface{}` - A pointer to an interface that will receive the decoded JSON data.

## Function Details

```go
package utils

import (
	"encoding/json"
	"os"
)

// ReadJsonFile reads a JSON file and unmarshals it into a Go interface.
func ReadJsonFile(path string, target interface{}) error {
	var content []byte

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, target)
	return err
}
```

## Usage

Here's a brief example of how to use the `ReadJsonFile` function:

```go
type Config struct {
	Port int `json:"port"`
}

func main() {
	config := &Config{}
	err := utils.ReadJsonFile("config.json", config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Port)
}
```

In this example, the `utils.ReadJsonFile` function is used to read a JSON file named "config.json". The decoded data is assigned to a `Config` struct pointer and any errors encountered during the process are logged and handled appropriately. If no errors occur, the `Port` value from the decoded JSON data can be accessed using the `config` variable.
