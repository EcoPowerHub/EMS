package pv

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/simonvetter/modbus"

	"github.com/EcoPowerHub/shared/pkg/io"
	"github.com/EcoPowerHub/shared/pkg/objects"
)

type Equipment struct {
	logger   zerolog.Logger
	mc       *modbus.ModbusClient
	state    objects.DriverState
	host     string
	readings readings
	lastRead time.Time
}

func (e *Equipment) AddOrRefreshData() error {
	var (
		err error
		p_w float32
	)
	// Read register 0
	p_w, err = e.mc.ReadFloat32(0, modbus.INPUT_REGISTER)
	if err != nil {
		e.logger.Error().Err(err).Msg("Cannot read register 0")
		e.state.Value = objects.EquipmentStateUnreachable
		return err
	}

	e.state.Value = objects.EquipmentStateOnline
	e.readings.p_w = float64(p_w)
	e.lastRead = time.Now()
	return nil
}

type readings struct {
	p_w float64
}

func New(host string) *Equipment {
	return &Equipment{
		state: objects.DriverState{Value: objects.EquipmentStateInit},
		host:  host,
	}
}

func (e *Equipment) Configure() (err error) {
	// Create modbus client
	e.mc, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL: fmt.Sprintf("tcp://%s", e.host),
	})
	if err != nil {
		e.state.Value = objects.EquipmentStateError
		return
	}
	// Open connection
	if err = e.mc.Open(); err != nil {
		e.state.Value = objects.EquipmentStateError
		return
	}
	e.state.Value = objects.EquipmentStateOnline
	return
}

func (e *Equipment) State() objects.DriverState {
	return e.state
}

func (e *Equipment) Read() map[string]map[string]any {
	return map[string]map[string]any{
		io.KeyPV: {
			io.KeyPV: objects.PV{
				P_kW:      e.readings.p_w / 1000.0,
				Timestamp: e.lastRead.UnixMicro(),
			},
		},
	}
}

func (e *Equipment) Write(_ map[string]map[string]any) error {
	return fmt.Errorf("Driver does not support writing")
}
