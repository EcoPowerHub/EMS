package ems

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/EcoPowerHub/EMS/config"
	manager "github.com/EcoPowerHub/EMS/driver"
	"github.com/EcoPowerHub/EMS/services"
	"github.com/EcoPowerHub/EMS/utils"
	context "github.com/EcoPowerHub/context/pkg"
	client "github.com/EcoPowerHub/triposter/pkg"
)

var ems core

type core struct {
	configuration    config.EMS
	equipmentManager *manager.Manager
	serviceManager   *services.Manager
	context          *context.Context
	triposter        client.Triposter
}

// Start reads the configuration
func Start(confpath string) {
	var (
		err error
	)

	err = utils.ParseConfFile(confpath, &ems.configuration)
	if err != nil {
		log.Error().Str("Error:", err.Error()).Msg("Reading Conf")
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
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to create context")
		return
	}

	// Create a managerEquipment with the parsed config
	managerEquipment, err := manager.New(ems.configuration.Equipments, ems.context)
	if err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to create managerEquipment")
		return
	}

	ems.equipmentManager = managerEquipment
	if err := ems.equipmentManager.SetupEquipments(); err != nil {
		log.Error().Str("Error:", err.Error()).Msg("Failed to setup equipments")
		return
	}

	servicesManager, err := services.New(ems.configuration.Services, ems.context, ems.configuration.Modes)
	if err != nil {
		log.Error().Str("Error:", err.Error()).Msg("Failed to create servicesManager")
		return
	}

	ems.serviceManager = servicesManager
	if err := ems.serviceManager.SetupServices(); err != nil {
		log.Error().Str("Error:", err.Error()).Msg("Failed to setup services")
		return
	}

	// Create a triposter client
	ems.triposter = client.New(ems.configuration.Triposter, ems.context, log.Logger)
	err = ems.triposter.Configure()
	if err != nil {
		log.Error().Str("Error:", err.Error()).Msg("Failed to configure triposter")
		return
	}

	go ems.triposter.Start()

	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := ems.equipmentManager.InitCycle(); err != nil {
			log.Error().Str("Error:", err.Error()).Msg("Failed to init cycle")
			return
		}

		// Reading drivers values
		err = ems.equipmentManager.Read()
		if err != nil {
			log.Error().Str("Error:", err.Error()).Msg("Failed to read")
			return
		}

		if err := ems.serviceManager.UpdateServices(); err != nil {
			log.Error().Str("Error:", err.Error()).Msg("Failed to update services")
			return
		}

		// Writing the context outputs to drivers
		err = ems.equipmentManager.Write()
		if err != nil {
			log.Error().Str("Error:", err.Error()).Msg("Failed to write")
			return
		}

		time.Sleep(1 * time.Second)
	}

}
