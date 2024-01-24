package config

import (
	context "github.com/EcoPowerHub/context/pkg"
)

/*
This package defines how is structured the complete config file
*/
type EMS struct {
	Equipments []Equipment           `json:"equipments"`
	Contexts   context.Configuration `json:"context"`
}
