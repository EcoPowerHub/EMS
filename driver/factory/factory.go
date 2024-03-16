package equipment

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/config"
	battery "github.com/EcoPowerHub/EMS/driver/drivers/generic/Battery"
	pv "github.com/EcoPowerHub/EMS/driver/drivers/generic/PV"
	poc "github.com/EcoPowerHub/EMS/driver/drivers/generic/Poc"
	"github.com/EcoPowerHub/shared/pkg/objects"
)

type Driver interface {
	Configure() error
	AddOrRefreshData() error
	State() objects.DriverState
	Read() map[string]map[string]any
	Write(map[string]map[string]any) error
}

type ManagerObject struct {
	Driver     Driver
	Equipement config.Equipment
}

// List equipments json to list Equipments
func Instanciate(listEquipmentsJson []config.Equipment) ([]ManagerObject, error) {

	equipments := make([]ManagerObject, len(listEquipmentsJson))

	for i, equipmentInList := range listEquipmentsJson {
		newEquipment, err := newDriver(equipmentInList)
		if err != nil {
			return nil, err
		}
		equipments[i].Driver, equipments[i].Equipement = newEquipment, equipmentInList
	}
	return equipments, nil
}

func newDriver(equipmentJson config.Equipment) (Driver, error) {
	switch equipmentJson.Id {
	case "generic/pv":
		return pv.New(equipmentJson.Host), nil
	case "generic/battery":
		return battery.New(equipmentJson.Host), nil
	case "generic/poc":
		return poc.New(equipmentJson.Host), nil
	default:
		return nil, fmt.Errorf("unsupported driver type: %s", equipmentJson.Id)
	}
}
