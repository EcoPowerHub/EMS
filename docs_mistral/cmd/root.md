 # EMS Package Documentation

This package contains the main entry point for the EMS (EcoPowerHub Energy Management System) application written in Go. It uses the `cobra` library for creating a CLI (Command Line Interface).

## Importing the package

```go
import (
	"log"
	"os"

	ems "github.com/EcoPowerHub/EMS/EMS"
	"github.com/spf13/cobra"
)
```

## Package Overview

The `cmd` package initializes and executes the EMS application using the provided configuration file. It also handles error handling and CLI arguments.

## Variables

### rootCmd

It is a pointer to the main cobra command that represents the application name "EMS". The `Use` field provides a short alias for this command, while the `Short` and `Long` fields give brief descriptions of the application.

```go
var (
	cfgFile string
	rootCmd = &cobra.Command{
		// Use: "EMS"...
	}
)
```

### cfgFile

It is a global variable to store the path to the configuration file, which will be read during application startup.

```go
func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "conf", "c", "", "(required) path to configuration file")
}
```

## Functions

### Execute

This function initializes and executes the EMS application. It checks if the `--conf` flag is provided or not before running the application. If it is missing, an error message is displayed and the application exits with an error code.

```go
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// Start EMS if no error
	ems.Start(cfgFile)
}
```

## Helper function: init()

The `init()` function sets the configuration file flag for the root command using cobra's PersistentFlags method. This makes it available to all commands in the application.

```go
func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "conf", "c", "", "(required) path to configuration file")
}
```
