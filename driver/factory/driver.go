package equipment

import "github.com/EcoPowerHub/EMS/common"

type Driver interface {
	Configure() error
	Cycle() error
	State() common.DriverState
	Read() map[string]map[string]any
	Write(map[string]map[string]any) error
}
