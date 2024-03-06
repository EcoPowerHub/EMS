Here's a Markdown documentation describing the functionality of your service:

---

# Evaluation Service Documentation

## Overview

The Evaluation Service is designed to dynamically evaluate expressions based on configurable inputs and outputs. It utilizes a generic evaluation library (`gval`) to process expressions and generate results, which are then used to update specified outputs in a context.

## Structure

The service consists of several key components:

- **Configuration (`conf`)**: Holds the settings for the service, including the expression to be evaluated.
- **Logger (`logger`)**: Utilized for logging service activities and errors.
- **Context (`ctx`)**: The execution context that provides access to external values and attributes necessary for input and output operations.
- **Inputs (`inputs`)**: A map defining the sources of values used in the expression evaluation.
- **Outputs (`outputs`)**: A map specifying where to store the results of the expression evaluation.
- **Values (`values`)**: A temporary storage for input values fetched from the context, used during expression evaluation.
- **Result (`result`)**: The final result of the expression evaluation, stored as a metric.

## Functions

### New

```go
func New(logger zerolog.Logger, ctx *context.Context) *service
```

Initializes a new instance of the service with the provided logger and context.

### Configure

```go
func (s *service) Configure(configuration any, inputsConf any, outputsConf any) error
```

Configures the service with the given settings, including the expression to evaluate, input sources, and output targets. The configuration is decoded and stored within the service instance for use during updates.

### Update

```go
func (s *service) Update() error
```

Performs the main logic of the service, which includes:

1. Updating input values by fetching the latest data from the context based on the configured inputs.
2. Evaluating the configured expression using the updated input values.
3. Converting the raw result of the expression into a metric.
4. Updating the configured outputs in the context with the result of the expression evaluation.

### updateInputs

```go
func (s *service) updateInputs() error
```

Fetches the latest values for each configured input from the context and stores them in the `values` map for use in expression evaluation.

### updateOutputs

```go
func (s *service) updateOutputs() error
```

Updates the configured outputs in the context with the result of the expression evaluation, effectively publishing the new values.

## Usage

To use the Evaluation Service, follow these steps:

1. **Initialization**: Create an instance of the service using the `New` function.
2. **Configuration**: Call `Configure` with the necessary configuration objects to set up the service.
3. **Execution**: Regularly invoke `Update` to evaluate the expression and update outputs based on the latest inputs.

The service dynamically evaluates expressions and updates outputs, making it versatile for applications requiring real-time data processing and decision-making.
