package driver

import (
	"fmt"
	"os"

	pv "github.com/EcoPowerHub/EMS/driver/drivers/generic/PV"
	"github.com/rs/zerolog"
)

type conf struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	Host        string      `json:"host"`
	DConf       interface{} `json:"conf"`
}

func newDriver(c conf, log zerolog.Logger) (Driver, error) {
	// Create corresponding logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Str("driver", c.Id).Logger()
	switch c.Id {
	case "generic/pv":
		return pv.New(logger, c.Host, c.DConf), nil
	default:
		return nil, fmt.Errorf("unsupported driver type: %s", c.Id)
	}
}
