package energytarget

import (
	"fmt"
	"time"

	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"go.einride.tech/pid"
)

type service struct {
	conf       conf
	logger     zerolog.Logger
	ctx        *context.Context
	controller *pid.Controller

	// inputs
	inputs inputs

	// outputs
	outputs outputs

	ess_soc *objects.Metric

	setpoint *objects.Setpoint

	period time.Duration
}

func New(logger zerolog.Logger, ctx *context.Context) *service {
	return &service{
		logger: logger,
		ctx:    ctx,
	}
}

func (s *service) Configure(configuration any, inputsConf any, outputsConf any) error {
	var (
		err error
	)
	if err := mapstructure.Decode(configuration, &s.conf); err != nil {
		return fmt.Errorf("invalid configuration: %v, error: %s", configuration, err)
	}

	if err := mapstructure.Decode(inputsConf, &s.inputs); err != nil {
		return fmt.Errorf("invalid inputs: %v", inputsConf)
	}

	if err := mapstructure.Decode(outputsConf, &s.outputs); err != nil {
		return fmt.Errorf("invalid outputs: %v", outputsConf)
	}

	if err := s.ValidateIO(); err != nil {
		return fmt.Errorf("invalid inputs or outputs: %s", err)
	}

	// Configure PID controller
	s.controller = &pid.Controller{
		Config: pid.ControllerConfig{
			ProportionalGain: s.conf.PID.Kp,
			IntegralGain:     s.conf.PID.Ki,
			DerivativeGain:   s.conf.PID.Kd,
		},
	}

	s.period, err = time.ParseDuration(s.conf.PID.Period)
	if err != nil {
		return fmt.Errorf("invalid period: %s", err)
	}

	return nil
}

func (s *service) Update() error {
	if err := s.updateInputs(); err != nil {
		return err
	}

	s.controller.Update(pid.ControllerInput{
		ReferenceSignal:  float64(s.conf.Target),
		ActualSignal:     s.ess_soc.Value,
		SamplingInterval: s.period,
	})

	s.setpoint = &objects.Setpoint{
		P_kW:      s.controller.State.ControlSignal,
		Timestamp: time.Now(),
	}

	return s.updateOutputs()
}

func (s *service) updateInputs() error {
	var (
		obj objects.Value
		err error
	)

	if obj, err = s.ctx.GetValue(s.inputs.Ess.Ref); err != nil {
		return err
	}

	//
	s.ess_soc, err = obj.GetAttr(s.inputs.Ess.Ref)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) updateOutputs() error {
	var (
		err error
	)
	// s.setpoint => p_kw
	if err = s.ctx.Set(s.outputs.Setpoint.Ref, s.setpoint); err != nil {
		return fmt.Errorf("failed to set setpoint: %s", err)
	}

	return nil
}

// ValidateIO is a function to ensure that the mandatory inputs and outputs are set
func (s *service) ValidateIO() error {
	if s.inputs.Ess.Ref == "" {
		return fmt.Errorf("Ess ref is not set")
	}
	if s.outputs.Setpoint.Ref == "" {
		return fmt.Errorf("setpoint is not set")
	}
	return nil
}
