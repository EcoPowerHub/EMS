package equipment

import (
	"fmt"
	"os"

	"github.com/EcoPowerHub/EMS/config"
	pv "github.com/EcoPowerHub/EMS/driver/drivers/generic/PV"
	"github.com/rs/zerolog"
)

func Instanciate(equipments config.Equipments) ([]Driver, error) {
	drivers := make([]Driver, len(equipments.Equipments))
	for i, equipment := range equipments.Equipments {
		// TODO replace logger instanciation, see #8
		driver, err := newDriver(equipment, zerolog.New(os.Stdout).With().Timestamp().Logger())
		if err != nil {
			return nil, err
		}
		drivers[i] = driver
	}
	return drivers, nil
}

func newDriver(c config.Equipment, log zerolog.Logger) (Driver, error) {
	// TODO replace logger instanciation, see #8
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Str("driver", c.Id).Logger()
	switch c.Id {
	case "generic/pv":
		return pv.New(logger, c.Host, c.DriverConf), nil
	default:
		return nil, fmt.Errorf("unsupported driver type: %s", c.Id)
	}
}
