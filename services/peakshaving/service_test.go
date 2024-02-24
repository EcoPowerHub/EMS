package peakshaving

import (
	"reflect"
	"testing"
)

func TestModule_Configure(t *testing.T) {
	// Create a new instance of the module
	m := &module{}
	period := "1s"

	// Define test cases
	tests := []struct {
		name           string
		configuration  any
		inputsConf     any
		outputsConf    any
		expectedError  bool
		expectedConfig conf
	}{
		{
			name: "Valid configuration",
			configuration: conf{
				PID: pidConf{
					Kp:     0.5,
					Ki:     0.1,
					Kd:     (0.2),
					Period: period,
				},
			},
			inputsConf: inputs{
				Poc_kW:   refWithAttr{Ref: "poc_kW", Attr: "value"},
				Limit_kW: refWithAttr{Ref: "limit_kW", Attr: "value"},
			},
			outputsConf: outputs{
				Setpoint: ref{Ref: "setpoint"},
			},
			expectedConfig: conf{
				PID: pidConf{
					Kp:     (0.5),
					Ki:     (0.1),
					Kd:     (0.2),
					Period: period,
				},
			},
			expectedError: false,
		},
		{
			name:          "Invalid configuration",
			configuration: "invalid",
			inputsConf: inputs{
				Poc_kW:   refWithAttr{Ref: "poc_kW", Attr: "value"},
				Limit_kW: refWithAttr{Ref: "limit_kW", Attr: "value"},
			},
			outputsConf: outputs{
				Setpoint: ref{Ref: "setpoint"},
			},
			expectedConfig: conf{},
			expectedError:  true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.Configure(tt.configuration, tt.inputsConf, tt.outputsConf)

			// Check if the error matches the expected error
			if (err == nil) == tt.expectedError {
				t.Errorf("unexpected error: got %v, want %v", err, tt.expectedError)
			}

			// Check if the configuration matches the expected configuration
			if !reflect.DeepEqual(m.conf, tt.expectedConfig) {
				t.Errorf("unexpected configuration: got %v, want %v", m.conf, tt.expectedConfig)
			}
		})
		// Reset the module
		m = new(module)
	}
}
func TestModule_ValidateIO(t *testing.T) {
	m := &module{
		inputs: inputs{
			Poc_kW:   refWithAttr{Ref: "poc_kW", Attr: "value"},
			Limit_kW: refWithAttr{Ref: "limit_kW", Attr: "value"},
		},
		outputs: outputs{
			Setpoint: ref{Ref: "setpoint"},
		},
	}

	tests := []struct {
		name           string
		inputs         inputs
		outputs        outputs
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name:          "Valid inputs and outputs",
			inputs:        m.inputs,
			outputs:       m.outputs,
			expectedError: false,
		},
		{
			name:           "Missing Poc_kW ref",
			inputs:         inputs{Poc_kW: refWithAttr{Attr: "value"}},
			outputs:        m.outputs,
			expectedError:  true,
			expectedErrMsg: "Poc_kW ref is not set",
		},
		{
			name:           "Missing Poc_kW attr",
			inputs:         inputs{Poc_kW: refWithAttr{Ref: "poc_kW"}},
			outputs:        m.outputs,
			expectedError:  true,
			expectedErrMsg: "Poc_kW attr is not set",
		},
		{
			name:           "Missing Limit_kW ref",
			inputs:         inputs{Poc_kW: refWithAttr{Ref: "poc_kW", Attr: "value"}},
			outputs:        m.outputs,
			expectedError:  true,
			expectedErrMsg: "Limit_kW ref is not set",
		},
		{
			name:           "Missing Limit_kW attr",
			inputs:         inputs{Poc_kW: refWithAttr{Ref: "poc_kW", Attr: "value"}, Limit_kW: refWithAttr{Ref: "limit_kW"}},
			outputs:        m.outputs,
			expectedError:  true,
			expectedErrMsg: "Limit_kW attr is not set",
		},
		{
			name:           "Missing setpoint",
			inputs:         m.inputs,
			outputs:        outputs{},
			expectedError:  true,
			expectedErrMsg: "setpoint is not set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.inputs = tt.inputs
			m.outputs = tt.outputs

			err := m.ValidateIO()

			if (err != nil) != tt.expectedError {
				t.Errorf("unexpected error: got %v, want %v", err, tt.expectedError)
			}

			if err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("unexpected error message: got %v, want %v", err.Error(), tt.expectedErrMsg)
			}
		})
	}
}
