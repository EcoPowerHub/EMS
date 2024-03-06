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

type core struct {
	configuration config.EMS
	manager       *manager.Manager
	context       *context.Context
	confpath      string
}

func New(confpath string) *core {
	return &core{
		confpath: confpath,
	}
}

func (c *core) configure() error {
	err := utils.ReadJsonFile(c.confpath, &c.configuration)
	if err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Reading Conf")
		return err
	}

	if !c.configuration.Debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	// Create the context
	c.context, err = context.New(c.configuration.Contexts)
	if err != nil {
		// #14
		fmt.Printf("Failed to create contexts: %s\n", err)
		return err
	}

	// Create a managerDriver with the parsed config
	managerDriver, err := manager.New(c.configuration.Drivers, c.context)
	if err != nil {
		// #8
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to create managerDriver")
		return err
	}

	c.manager = managerDriver
	if err := c.manager.SetupDrivers(); err != nil {
		log.Fatal().Str("Error:", err.Error()).Msg("Failed to setup drivers")
		return err
	}

	return nil
}

// Start reads the configuration
func (c *core) Start() {
	var (
		err error
	)

	if c.configure() != nil {
		return
	}
	// While cycle isn't finished
	for {
		// Executing all drivers
		if err := c.manager.InitCycle(); err != nil {
			return
		}

		// Reading drivers values
		err = c.manager.Read()
		if err != nil {
			fmt.Println(err)
		}

		// Writing the context outputs to drivers
		err = c.manager.Write()
		fmt.Printf("Readings %s\n", c.context)
		fmt.Printf("Writings error %s\n", err)
		time.Sleep(1 * time.Second)
	}

}
