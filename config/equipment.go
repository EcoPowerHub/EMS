package config

/*
This package defines how is structured the equipment part in the configuration file.
*/
type Equipment struct {
	Id          string                    `json:"id"`
	Description string                    `json:"description"`
	Host        string                    `json:"host"`
	Name        string                    `json:"name"`
	DriverConf  interface{}               `json:"conf"`
	Inputs      map[string]map[string]ref `json:"inputs"`
	Outputs     map[string]map[string]ref `json:"outputs"`
}

type ref struct {
	Ref string `json:"ref"`
}
