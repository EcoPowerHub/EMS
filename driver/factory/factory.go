package equipment

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/config"
	pv "github.com/EcoPowerHub/EMS/driver/drivers/generic/PV"
)

type Driver interface {
	Configure() error
	AddOrRefreshData() error
	State() common.EquipmentState
	Read() map[string]map[string]any
	Write(map[string]map[string]any) error
}

// List equipments json to list Equipments
func Instanciate(listEquipmentsJson []config.Equipment) ([]Driver, error) {

	equipments := make([]Driver, len(listEquipmentsJson))

	for i, equipmentInList := range listEquipmentsJson {
		newEquipment, err := newDriver(equipmentInList)
		if err != nil {
			return nil, err
		}
		equipments[i] = newEquipment
	}
	return equipments, nil
}

func newDriver(equipmentJson config.Equipment) (Driver, error) {
	switch equipmentJson.Id {
	case "generic/pv":
		return pv.New(equipmentJson.Host), nil
	default:
		return nil, fmt.Errorf("unsupported driver type: %s", equipmentJson.Id)
	}
}
