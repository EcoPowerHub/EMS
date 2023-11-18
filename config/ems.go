package config

/*
This package defines how is structured the complete config file
*/

type EMS struct {
	Equipments []Equipment `json:"equipments"`
}
