package pv

import (
	"fmt"
	"time"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/common/io"
	"github.com/rs/zerolog"
	"github.com/simonvetter/modbus"
)

type Equipment struct {
	logger   zerolog.Logger
	mc       *modbus.ModbusClient
	state    common.EquipmentState
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
		e.state.Value = common.EquipmentStateUnreachable
		return err
	}

	e.state.Value = common.EquipmentStateOnline
	e.readings.p_w = float64(p_w)
	e.lastRead = time.Now()
	return nil
}

type readings struct {
	p_w float64
}

func New(host string) *Equipment {
	return &Equipment{
		state: common.EquipmentState{Value: common.EquipmentStateInit},
		host:  host,
	}
}

func (e *Equipment) Configure() (err error) {
	// Create modbus client
	e.mc, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL: fmt.Sprintf("tcp://%s", e.host),
	})
	if err != nil {
		e.state.Value = common.EquipmentStateError
		return
	}
	// Open connection
	if err = e.mc.Open(); err != nil {
		e.state.Value = common.EquipmentStateError
		return
	}
	e.state.Value = common.EquipmentStateOnline
	return
}

func (e *Equipment) State() common.EquipmentState {
	return e.state
}

func (e *Equipment) Read() map[string]map[string]any {
	return map[string]map[string]any{
		io.PV: {
			io.PV: common.PV{
				P_kW:      e.readings.p_w / 1000.0,
				Timestamp: e.lastRead.UnixMicro(),
			},
		},
	}
}

func (e *Equipment) Write(_ map[string]map[string]any) error {
	return nil
}
