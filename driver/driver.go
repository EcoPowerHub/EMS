package main

import (
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

type Driver interface {
	Configure()
	Start()
	Cycle()
}

type conf struct {
	host string `json:"host"`
}

func Configure(c conf) {
	var client *modbus.ModbusClient
	var err error

	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://" + c.host,
		Timeout: 1 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.ReadUint32(4, modbus.INPUT_REGISTER)

/*
func Start() {

}

func Cycle() None {

}
*/
