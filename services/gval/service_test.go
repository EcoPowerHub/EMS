package eval

import (
	"strings"
	"testing"
)

func TestConfigure(t *testing.T) {
	tests := []struct {
		name          string
		configuration conf
		inputsConf    map[string]refWithAttr
		outputsConf   map[string]refWithAttr
		expectError   bool
		errorMessage  string
	}{
		{
			name: "Valid configuration",
			configuration: conf{
				Expression: "2+2",
			},
			inputsConf: map[string]refWithAttr{
				"input1": {Ref: "sensor1", Attr: "temperature"},
			},

			outputsConf: map[string]refWithAttr{
				"output1": {Ref: "actuator1", Attr: "state"},
			},

			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}

			err := s.Configure(tt.configuration, tt.inputsConf, tt.outputsConf)

			if tt.expectError {
				if err == nil || !strings.Contains(err.Error(), tt.errorMessage) {
					t.Errorf("%s: Expected error to contain '%s', but got '%v'", tt.name, tt.errorMessage, err)
				}
			} else if err != nil {
				t.Errorf("%s: Expected no error, but got %v", tt.name, err)
			}
		})
	}
}
