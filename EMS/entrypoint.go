package ems

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

	err = utils.ReadJsonFile(confpath, &ems.configuration)
	if err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Reading Conf")
		return
	}

	if !ems.configuration.Debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	// Create the context
	ems.context, err = context.New(ems.configuration.Contexts)
	if err != nil {
		// #14
		fmt.Printf("Failed to create contexts: %s\n", err)
		return
	}

	// Create a managerEquipment with the parsed config
	managerEquipment, err := manager.New(ems.configuration.Equipments, ems.context)
	if err != nil {
		// #8
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to create managerEquipment")
		return
	}

	ems.manager = managerEquipment
	if err := ems.manager.SetupEquipments(); err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to setup equipments")
		return
	}

	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := ems.manager.InitCycle(); err != nil {
			return
		}

		// Reading drivers values
		err = ems.manager.Read()
		if err != nil {
			fmt.Println(err)
		}

		// Writing the context outputs to drivers
		err = managerEquipment.Write()

		fmt.Printf("Readings %s\n", ems.context)
		fmt.Printf("Writings error %s\n", err)
		time.Sleep(1 * time.Second)
	}

}
