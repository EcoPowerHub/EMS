package ems

import (
	"bytes"
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

	// Create a managerEquipment with the parsed config
	managerEquipment, err := manager.New(ems.configuration.Equipments)
	if err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to create managerEquipment")
		return
	}

	// Create the context
	ems.context, err = context.New(ems.configuration.Contexts)
	if err != nil {
		// #14
		fmt.Printf("Failed to create contexts: %s\n", err)
		return
	}

	ems.manager = managerEquipment
	if err := managerEquipment.SetupEquipments(); err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to setup equipments")
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
		log.Trace().Str("Reading", createKeyValuePairs(readings)).Msg("Readings")
		time.Sleep(1 * time.Second)
	}

}

// Convert map to string for log purposes
func createKeyValuePairs(m map[string]map[string]any) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}
