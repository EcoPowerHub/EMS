package ems

import (
	"fmt"
	"log"
	"time"

	"github.com/EcoPowerHub/EMS/config"
	manager "github.com/EcoPowerHub/EMS/driver"
	"github.com/EcoPowerHub/EMS/utils"
	context "github.com/EcoPowerHub/context/pkg"
)

var ems core

type core struct {
	configuration config.EMS
	manager       *manager.Manager
	context       *context.Context
}

// Start reads the configuration
func Start(confpath string) {
	var (
		err error
	)

	// Create the context
	context.New(ems.configuration.Contexts)

	err = utils.ReadJsonFile(confpath, &ems.configuration)
	if err != nil {
		// #8
		log.Fatalf("Error: %s\n", err)
		return
	}

	// Create a managerEquipment with the parsed config
	managerEquipment, err := manager.New(ems.configuration.Equipments)
	if err != nil {
		// #8
		fmt.Printf("Failed to create managerEquipment: %s\n", err)
		return
	}

	ems.manager = managerEquipment
	if err := managerEquipment.SetupEquipments(); err != nil {
		// #8
		fmt.Printf("Failed to setup equipments: %s\n", err)
		return
	}

	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := managerEquipment.InitCycle(); err != nil {
			return
		}
		// Reading drivers values
		readings := managerEquipment.Read()
		fmt.Printf("Readings %s\n", readings)
		bla, _ := ems.context.Get("ess1")
		fmt.Printf("%s\n", bla)
		time.Sleep(1 * time.Second)
	}

}
