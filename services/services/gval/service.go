package eval

import (
	"fmt"
	"reflect"

	eval "github.com/EcoPowerHub/EMS/helpers/gval"
	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
	"github.com/mitchellh/mapstructure"

	"github.com/rs/zerolog"
)

type service struct {
	conf   conf
	logger zerolog.Logger
	ctx    *context.Context

	inputs  map[string]refWithAttr
	outputs map[string]refWithAttr

	values map[string]any

	result *objects.Metric
}

func New(logger zerolog.Logger, ctx *context.Context) *service {
	return &service{
		logger: logger,
		ctx:    ctx,
		values: make(map[string]any),
	}
}

func (s *service) Configure(configuration any, inputsConf any, outputsConf any) error {
	var (
		err error
	)

	if err = mapstructure.Decode(configuration, &s.conf); err != nil {
		return fmt.Errorf("failed to decode configuration: %w", err)
	}

	if err = mapstructure.Decode(inputsConf, &s.inputs); err != nil {
		return fmt.Errorf("failed to decode inputs configuration: %w", err)
	}

	if err = mapstructure.Decode(outputsConf, &s.outputs); err != nil {
		return fmt.Errorf("failed to decode outputs configuration: %w", err)
	}

	return nil
}

func (s *service) Update() error {
	var (
		rawResult any
		err       error
		kind      reflect.Kind
	)

	if err = s.updateInputs(); err != nil {
		return fmt.Errorf("failed to update inputs: %w", err)
	}

	rawResult, kind, err = eval.Evaluate(s.conf.Expression, s.values)
	if err != nil {
		return fmt.Errorf("failed to evaluate expression: %w", err)
	}

	s.result, err = eval.ToMetric(rawResult, kind)
	if err != nil {
		return fmt.Errorf("failed to convert result to metric: %w", err)
	}

	if err = s.updateOutputs(); err != nil {
		return fmt.Errorf("failed to update outputs: %w", err)
	}

	return nil
}

func (s *service) updateInputs() error {
	var (
		ref objects.Value
		val *objects.Metric
		err error
	)
	for name, input := range s.inputs {

		ref, err = s.ctx.GetValue(input.Ref)
		if err != nil {
			return fmt.Errorf("failed to get object for %s: %w", name, err)
		}

		if val, err = ref.GetAttr(input.Attr); err != nil {
			return fmt.Errorf("failed to get attribute for %s: %w", name, err)
		}

		s.values[name] = val.Value
	}
	return nil
}

func (s *service) updateOutputs() error {
	var (
		ref objects.Value
		err error
	)

	for name, output := range s.outputs {
		ref, err = s.ctx.GetValue(output.Ref)
		if err != nil {
			return fmt.Errorf("failed to get object for %s: %w", name, err)
		}

		if err = ref.SetAttr(output.Attr, s.result.Value); err != nil {
			return fmt.Errorf("failed to set attribute for %s: %w", name, err)
		}
	}
	return nil
}
