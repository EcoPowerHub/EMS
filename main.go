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
	fmt.Println("Starting app")
	var equipmentData config.Equipments
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	err := utils.ReadJsonFile("tests/equipment.json", &equipmentData)
	if err != nil {
		return
	}
	equipmentData.Print()

	manager, err := manager.New(equipmentData)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create manager")
		return
	}
	if err := manager.Configure(); err != nil {
		logger.Error().Err(err).Msg("Failed to configure manager")
		return
	}

	for {
		if err := manager.Cycle(); err != nil {
			return
		}
		readings := manager.Read()
		logger.Info().Interface("readings", readings).Msg("Readings")

		time.Sleep(1 * time.Second)
	}
}
