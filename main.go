package main

import (
	"fmt"
	"os"
	"time"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/driver/manager"
	"github.com/EcoPowerHub/EMS/utils"
	"github.com/rs/zerolog"
)

func main() {

	// Create logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	fmt.Println("Starting app")

	var equipmentData config.Equipments // []Equipment

	// Read the file and make it matching with equipmentData
	err := utils.ReadJsonFile("tests/equipment.json", &equipmentData)
	if err != nil {
		return
	}

	// Print the equipment configuration
	equipmentData.Print()

	// Create a manager (drivers) with the config
	manager, err := manager.New(equipmentData)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create manager")
		return
	}
	if err := manager.Configure(); err != nil {
		logger.Error().Err(err).Msg("Failed to configure manager")
		return
	}

	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := manager.Cycle(); err != nil {
			return
		}
		// Reading drivers values
		readings := manager.Read()
		logger.Info().Interface("readings", readings).Msg("Readings")

		time.Sleep(1 * time.Second)
	}
}
