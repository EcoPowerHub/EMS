// battery.go
package main

import (
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

// conf doit être exporté, alors changez son nom à Conf.
type Conf struct {
	Host string `json:"host"`
}

// Battery implémente l'interface Driver
type Battery struct {
	Config Conf
	Client *modbus.ModbusClient
}

// Configure implémente la méthode Configure de l'interface Driver
func (b *Battery) Configure(c Conf) {
	b.Config = c
	var err error
	b.Client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://" + c.Host,
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	b.Client.ReadUint32(4, modbus.INPUT_REGISTER)
}

// Start implémente la méthode Start de l'interface Driver
func (b *Battery) Start() {
	// Implémentation de la logique pour la méthode Start
}

// Cycle implémente la méthode Cycle de l'interface Driver
func (b *Battery) Cycle() {
	// Implémentation de la logique pour la méthode Cycle
}
