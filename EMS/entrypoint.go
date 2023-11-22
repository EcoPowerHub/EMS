package ems

import (
	"log"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/utils"
)

func Start(confpath string) {
	var (
		configuration config.EMS
		err           error
	)

	err = utils.ReadJsonFile(confpath, &configuration)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		return
	}

	log.Printf("Configuration file: %s\n", confpath)

}
