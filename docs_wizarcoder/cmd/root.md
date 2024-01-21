```go
package cmd

import (
	"log"
	"os"

	"github.com/EcoPowerHub/EMS/EMS"
	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd = &cobra.Command{
    Use: "EMS", // (To define)
    Short: "(To define)", // (To define)
    Long: `long description (to define)`,
    Run: func(cmd *cobra.Command, args []string) {
        if cfgFile == "" {
            // #8
            log.Fatal("Missing required flag: --conf.")
            _ = cmd.Help()
            os.Exit(1)
            return
        }
    },
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }

    // Start EMS if no error
    EMS.Start(cfgFile)
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&cfgFile, "conf", "c", "", "(required) path to configuration file")
}
```
