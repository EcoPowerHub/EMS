Package config
```go
// Package config defines the configuration for an Equipment.
package config

import "encoding/json"

type Equipment struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	Host        string      `json:"host"`
	Name       string      `json:"name"`
	DriverConf interface{} `json:"conf"`
}
```

This package defines the structure of an Equipment in a configuration file. An Equipment is a struct that contains four fields, Id, Description, Host and Name, which are all strings. The DriverConf field can be any type of data as it is defined as an interface{}. This means that it can hold any value and will be marshaled/unmarshaled according to its actual value when serializing and deserializing the configuration. For example:

```go
type Equipment struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	Host        string      `json:"host"`
	Name       string      `json:"name"`
	DriverConf  interface{} `json:"conf"`
}

type EquipmentWithJson struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	Host        string      `json:"host"`
	Name       string      `json:"name"`
	DriverConf  json.RawMessage `json:"conf"`
}
```

This configuration can hold a JSON object with any number of key-value pairs as long as they can be represented by an interface{} type. This allows great flexibility and versatility in configuring the equipment's driver. Here is an example:

```go
{
	"id": "eq1",
	"description": "Equipment Description 1",
	"host": "https://example.com",
	"name": "equipment-name",
	"conf": {
		"key1": "value1",
		"key2": true,
		"key3": 10,
		"key4": [
			"item1",
			"item2"
	]
}
```
