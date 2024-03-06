package peakshaving

import (
	"fmt"
	"time"

	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"go.einride.tech/pid"
)

type module struct {
	conf       conf
	logger     zerolog.Logger
	ctx        *context.Context
	controller *pid.Controller

	// inputs
	inputs inputs

	// outputs
	outputs outputs

	poc_kW   *objects.Metric
	limit_kW *objects.Metric

	setpoint *objects.Setpoint

	period time.Duration
}

func New(logger zerolog.Logger, ctx *context.Context) *module {
	return &module{
		logger: logger,
		ctx:    ctx,
	}
}

func (m *module) Configure(configuration any, inputsConf any, outputsConf any) error {
	var (
		err error
	)
	if err := mapstructure.Decode(configuration, &m.conf); err != nil {
		return fmt.Errorf("invalid configuration: %v, error: %s", configuration, err)
	}

	if err := mapstructure.Decode(inputsConf, &m.inputs); err != nil {
		return fmt.Errorf("invalid inputs: %v", inputsConf)
	}

	if err := mapstructure.Decode(outputsConf, &m.outputs); err != nil {
		return fmt.Errorf("invalid outputs: %v", outputsConf)
	}

	if err := m.ValidateIO(); err != nil {
		return fmt.Errorf("invalid inputs or outputs: %s", err)
	}

	// Configure PID controller
	m.controller = &pid.Controller{
		Config: pid.ControllerConfig{
			ProportionalGain: m.conf.PID.Kp,
			IntegralGain:     m.conf.PID.Ki,
			DerivativeGain:   m.conf.PID.Kd,
		},
	}

	m.period, err = time.ParseDuration(m.conf.PID.Period)
	if err != nil {
		return fmt.Errorf("invalid period: %s", err)
	}

	return nil
}

func (m *module) Update() error {
	if err := m.updateInputs(); err != nil {
		return err
	}

	m.controller.Update(pid.ControllerInput{
		ReferenceSignal:  m.limit_kW.Value,
		ActualSignal:     m.poc_kW.Value,
		SamplingInterval: m.period,
	})

	m.setpoint = &objects.Setpoint{
		P_kW:      m.controller.State.ControlSignal,
		Timestamp: time.Now(),
	}

	return m.updateOutputs()
}

func (m *module) updateInputs() error {
	var (
		obj objects.Value
		err error
	)

	if obj, err = m.ctx.GetValue(m.inputs.Poc_kW.Ref); err != nil {
		return err
	}

	m.poc_kW, err = obj.GetAttr(m.inputs.Poc_kW.Attr)
	if err != nil {
		return err
	}

	if obj, err = m.ctx.GetValue(m.inputs.Limit_kW.Ref); err != nil {
		return err
	}

	m.limit_kW, err = obj.GetAttr(m.inputs.Limit_kW.Attr)
	if err != nil {
		return err
	}

	return nil
}

func (m *module) updateOutputs() error {
	var (
		err error
	)
	if err = m.ctx.Set(m.outputs.Setpoint.Ref, m.setpoint); err != nil {
		return fmt.Errorf("failed to set setpoint: %s", err)
	}

	if m.outputs.PidError.Ref != "" {
		if err = m.ctx.Set(m.outputs.PidError.Ref, m.controller.State.ControlError); err != nil {
			return fmt.Errorf("failed to set pid error: %s", err)
		}
	}

	return nil
}

// ValidateIO is a function to ensure that the mandatory inputs and outputs are set
func (m *module) ValidateIO() error {
	if m.inputs.Poc_kW.Ref == "" {
		return fmt.Errorf("Poc_kW ref is not set")
	}
	if m.inputs.Poc_kW.Attr == "" {
		return fmt.Errorf("Poc_kW attr is not set")
	}
	if m.inputs.Limit_kW.Ref == "" {
		return fmt.Errorf("Limit_kW ref is not set")
	}
	if m.inputs.Limit_kW.Attr == "" {
		return fmt.Errorf("Limit_kW attr is not set")
	}
	if m.outputs.Setpoint.Ref == "" {
		return fmt.Errorf("setpoint is not set")
	}
	return nil
}
