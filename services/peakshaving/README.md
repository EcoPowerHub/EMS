# Peak Shaving Module Documentation

## Overview

The Peak Shaving Module is designed to optimize energy consumption by managing and adjusting power usage to avoid peak demand charges. This module utilizes a PID (Proportional, Integral, Derivative) controller to dynamically adjust setpoints based on current and limit power values.

## How to Use the Peak Shaving Module

To utilize the Peak Shaving Module effectively, you need to configure it properly within your service architecture. Below is a step-by-step guide on how to set up and use the Peak Shaving Module, including an example configuration.

###  Define the Configuration

Your configuration should be structured to include the module's PID controller settings, input references, and output references. Here's an example of how the configuration might look in a JSON format:

```json
"services": {
  "peakshaving": {
    "id": "peakshaving",
    "priority": 1,
    "conf": {
      "pid_controller": {
        "kp": 0.1,
        "ki": 0.1,
        "kd": 0.1,
        "period": "1s"
      },
      "inputs": {
        "poc_kW": {
          "ref": "poc",
          "attr": "value"
        },
        "limit_kW": {
          "ref": "limit",
          "attr": "value"
        }
      },
      "outputs": {
        "setpoint": {
          "ref": "setpoint"
        },
        "pid_error": {
          "ref": "pidError" // Optional: Only if you want to monitor the PID error
        }
      }
    }
  }
}
```


## Components

### Configuration

- **PID Settings**: Configures the PID controller with `ProportionalGain`, `IntegralGain`, `DerivativeGain`, and the control `Period`.
- **Inputs & Outputs**: Defines the references and attributes for input metrics (`Poc_kW`, `Limit_kW`) and output setpoints (`Setpoint`), including any PID error outputs.

### Methods

#### `New(logger zerolog.Logger, ctx *context.Context) *module`

Constructor function that initializes a new instance of the peak shaving module with the provided logger and context.

#### `Configure(configuration any, inputsConf any, outputsConf any) error`

Configures the module with specified configuration settings for the module itself, inputs, and outputs. This method also initializes the PID controller based on the configuration.

#### `Update() error`

Performs an update cycle of the module, fetching the latest input values, calculating the new setpoint via the PID controller, and updating the outputs accordingly.

#### `ValidateIO() error`

Validates that all required inputs and outputs are properly configured. This is called during the `Configure` method to ensure mandatory configurations are set.


## Error Handling

Errors are returned by each method to indicate issues in configuration, input/output validation, or during the update cycle. Proper error handling should be implemented by the caller to ensure smooth operation.
