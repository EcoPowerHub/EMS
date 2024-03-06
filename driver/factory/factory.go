package driver

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/config"
	pv "github.com/EcoPowerHub/EMS/driver/drivers/generic/PV"
	"github.com/EcoPowerHub/shared/pkg/objects"
)

type Driver interface {
	Configure() error
	AddOrRefreshData() error
	State() objects.DriverState
	Read() map[string]map[string]any
	Write(map[string]map[string]any) error
}

type Manager struct {
	Drivers []Driver
	Conf    config.Driver
}

// List drivers json to list Drivers
func (m *Manager) Instanciate(listDriverJson []config.Driver) ([]Driver, error) {

	drivers := make([]Driver, len(listDriverJson))

	for _, driverInList := range listDriverJson {
		newDriver, err := newDriver(driverInList)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, newDriver)
	}

	return drivers, nil
}

func newDriver(driverJson config.Driver) (Driver, error) {
	switch driverJson.Id {
	case "generic/pv":
		return pv.New(driverJson.Host), nil
	default:
		return nil, fmt.Errorf("unsupported driver type: %s", driverJson.Id)
	}
}
