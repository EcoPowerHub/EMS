package config

import (
	"github.com/EcoPowerHub/EMS/services/modes"
	context "github.com/EcoPowerHub/context/pkg"
	triposter "github.com/EcoPowerHub/triposter/pkg"
)

/*
This package defines how is structured the complete config file
*/
type EMS struct {
	Equipments []Equipment             `json:"equipments"`
	Contexts   context.Configuration   `json:"context"`
	Services   map[string]Service      `json:"services"`
	Modes      modes.Conf              `json:"modes"`
	Debug      bool                    `json:"debug"`
	Triposter  triposter.Configuration `json:"triposter"`
}
