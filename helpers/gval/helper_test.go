package eval

import (
	"math"
	"reflect"
	"testing"

	"github.com/EcoPowerHub/shared/pkg/objects"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		values     map[string]any
		want       any
		wantKind   reflect.Kind
		wantErr    bool
	}{
		{
			name:       "Evaluate with variables success",
			expression: "2 + 3 * 2",
			values:     nil,
			want:       float64(8),
			wantKind:   reflect.Float64,
			wantErr:    false,
		},
		{
			name:       "Evaluate with variables and context",
			expression: "x + y",
			values:     map[string]any{"x": 2, "y": 3},
			want:       float64(5),
			wantKind:   reflect.Float64,
			wantErr:    false,
		},
		{
			name:       "Evaluate with invalid expression",
			expression: "2 +",
			values:     nil,
			want:       nil,
			wantKind:   reflect.Invalid,
			wantErr:    true,
		},
		{
			name:       "Empty expression",
			expression: "",
			values:     nil,
			want:       nil,
			wantKind:   reflect.Invalid,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotKind, err := Evaluate(tt.expression, tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() got = %v, want %v", got, tt.want)
			}
			if gotKind != tt.wantKind {
				t.Errorf("Evaluate() gotKind = %v, want %v", gotKind, tt.wantKind)
			}
		})
	}
}

func TestValidExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantValid  bool
	}{
		{
			name:       "Valid expression",
			expression: "2 + 3 * 2",
			wantValid:  true,
		},
		{
			name:       "Invalid expression",
			expression: "2 + 3 *",
			wantValid:  false,
		},
		{
			name:       "Empty expression",
			expression: "",
			wantValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, _ := ValidExpression(tt.expression)
			if valid != tt.wantValid {
				t.Errorf("validateExpression() got = %v, want %v for expression: %s", valid, tt.wantValid, tt.expression)
			}
		})
	}
}

func TestToMetric(t *testing.T) {
	tests := []struct {
		name    string
		value   any
		kind    reflect.Kind
		want    *objects.Metric
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Convert float64",
			value:   float64(3.14),
			kind:    reflect.Float64,
			want:    &objects.Metric{Value: 3.14},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Convert float32",
			value:   float32(2.71828),
			kind:    reflect.Float32,
			want:    &objects.Metric{Value: 2.71828},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Convert int",
			value:   42,
			kind:    reflect.Int,
			want:    &objects.Metric{Value: 42},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Convert uint",
			value:   uint(123),
			kind:    reflect.Uint,
			want:    &objects.Metric{Value: 123},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Convert bool (true)",
			value:   true,
			kind:    reflect.Bool,
			want:    &objects.Metric{Value: 1.0},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Convert bool (false)",
			value:   false,
			kind:    reflect.Bool,
			want:    &objects.Metric{Value: 0.0},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Unsupported type",
			value:   "unsupported",
			kind:    reflect.String,
			want:    nil,
			wantErr: true,
			errMsg:  "unsupported type for Metric conversion: string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMetric(tt.value, tt.kind)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMetric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("ToMetric() error message = %v, want %v", err.Error(), tt.errMsg)
				return
			}
			if tt.want != nil && got != nil && !almostEqual(got.Value, tt.want.Value, 0.0001) {
				t.Errorf("ToMetric() got = %v, want %v", got.Value, tt.want.Value)
			}
		})
	}
}

func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
