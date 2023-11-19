package pv

import (
	"fmt"
	"time"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/common/io"
	"github.com/rs/zerolog"
	"github.com/simonvetter/modbus"
)

type Driver struct {
	logger   zerolog.Logger
	mc       *modbus.ModbusClient
	state    common.DriverState
	host     string
	readings readings
	lastRead time.Time
}

type readings struct {
	p_w float64
}

func New(logger zerolog.Logger, host string, conf interface{}) *Driver {
	return &Driver{
		logger: logger,
		state:  common.DriverState{Value: common.DriverStateInit},
		host:   host,
	}
}

func (d *Driver) Configure() (err error) {

	// Create modbus client
	d.mc, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL: fmt.Sprintf("tcp://%s", d.host),
	})

	if err != nil {
		d.state.Value = common.DriverStateError
		return
	}

	// Open connection
	if err = d.mc.Open(); err != nil {
		d.state.Value = common.DriverStateError
		return
	}

	d.state.Value = common.DriverStateOnline
	return
}

func (d *Driver) State() common.DriverState {
	return d.state
}

func (d *Driver) Cycle() error {
	var (
		err error
		p_w float32
	)

	// Read register 0
	p_w, err = d.mc.ReadFloat32(0, modbus.INPUT_REGISTER)
	if err != nil {
		d.logger.Error().Err(err).Msg("Cannot read register 0")
		d.state.Value = common.DriverStateUnreachable
		return err
	}

	d.state.Value = common.DriverStateOnline
	d.readings.p_w = float64(p_w)
	d.lastRead = time.Now()
	return nil
}

func (d *Driver) Read() map[string]map[string]any {
	return map[string]map[string]any{
		io.PV: {
			io.PV: common.PV{
				P_kW:      d.readings.p_w / 1000.0,
				Timestamp: d.lastRead.UnixMicro(),
			},
		},
	}
}

func (d *Driver) Write(_ map[string]map[string]any) error {
	return nil
}
