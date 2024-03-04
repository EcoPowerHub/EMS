package eval

import (
	"fmt"
	"reflect"

	"github.com/EcoPowerHub/shared/pkg/objects"
	"github.com/PaesslerAG/gval"
)

func Evaluate(expression string, values map[string]any) (result any, kind reflect.Kind, err error) {
	if expression == "" {
		return nil, reflect.Invalid, nil
	}

	result, err = gval.Evaluate(expression, values)
	if err != nil {
		return nil, reflect.Invalid, err
	}

	if result == nil {
		return nil, reflect.Invalid, nil
	}

	return result, reflect.TypeOf(result).Kind(), nil
}

func ValidExpression(expression string) (bool, error) {
	_, err := gval.Evaluate(expression, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ToMetric(value any, kind reflect.Kind) (*objects.Metric, error) {
	switch kind {
	case reflect.Float64:
		return &objects.Metric{
			Value: value.(float64),
		}, nil
	case reflect.Float32:
		return &objects.Metric{
			Value: float64(value.(float32)),
		}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &objects.Metric{
			Value: float64(reflect.ValueOf(value).Int()),
		}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &objects.Metric{
			Value: float64(reflect.ValueOf(value).Uint()),
		}, nil
	case reflect.Bool:
		if value.(bool) {
			return &objects.Metric{
				Value: 1.0,
			}, nil
		}
		return &objects.Metric{
			Value: 0.0,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported type for Metric conversion: %v", kind.String())
	}
}
