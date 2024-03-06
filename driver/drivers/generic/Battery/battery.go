package battery

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
		singleAdrr []uint16
		err        error
		multiAdrr  []float32
	)

	// read register 0 and 1
	singleAdrr, err = e.mc.ReadRegisters(0, 2, modbus.INPUT_REGISTER)
	if err != nil {
		e.state.Value = objects.DriverStateUnreachable
		return fmt.Errorf("Unable to read soc and soh: ", err)
	}

	// read register 2 and 4
	multiAdrr, err = e.mc.ReadFloat32s(2, 2, modbus.INPUT_REGISTER)
	if err != nil {
		e.state.Value = objects.DriverStateUnreachable
		return fmt.Errorf("Unable to read capacity_Wh and Active Power: ", err)
	}

	e.readings = readings{
		SoC:         float64(singleAdrr[0]),
		SoH:         float64(singleAdrr[1]),
		Capacity_Wh: float64(multiAdrr[0] / 1000.0),
		P_W:         float64(multiAdrr[1]),
	}
	e.lastRead = time.Now()
	return nil
}

type readings struct {
	SoC         float64
	SoH         float64
	Capacity_Wh float64
	P_W         float64
}

func New(host string) *Equipment {
	return &Equipment{
		state: objects.DriverState{Value: objects.DriverStateInit},
		host:  host,
	}
}

func (e *Equipment) Configure() (err error) {
	// Create modbus client
	e.mc, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL: fmt.Sprintf("tcp://%s", e.host),
	})
	if err != nil {
		e.state.Value = objects.DriverStateError
		return
	}
	// Open connection
	if err = e.mc.Open(); err != nil {
		e.state.Value = objects.DriverStateError
		return err
	}
	e.state.Value = objects.DriverStateOnline
	return err
}

func (e *Equipment) State() objects.DriverState {
	return e.state
}

func (e *Equipment) Read() map[string]map[string]any {
	return map[string]map[string]any{
		io.KeyBattery: {
			io.KeyBattery: &objects.Battery{
				Timestamp:    time.Now(),
				SoC:          e.readings.SoC,
				SoH:          e.readings.SoH,
				Capacity_kWh: e.readings.Capacity_Wh / 1000.0,
				P_kW:         e.readings.P_W / 1000.0,
			},
		},
	}
}

func (e *Equipment) Write(writings map[string]map[string]any) error {
	setpoint, ok := writings[io.KeySetpoint][io.KeySetpoint]
	if !ok {
		return fmt.Errorf("Unable to get setpoint.")
	}
	a := setpoint.(*objects.Setpoint)

	err := e.mc.WriteFloat32(0, float32(a.P_kW)*1000.0)
	if err != nil {
		return fmt.Errorf("Error while writing the setpoint ", err)
	}
	return nil
}
