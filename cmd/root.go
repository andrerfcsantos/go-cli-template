package cmd

import (
	configCmd "app/cmd/config"
	"app/lib/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const appName = "myapp"

var cfg *config.Config

func init() {
	cfg = config.NewConfig(appName)
	rootCmd.AddCommand(configCmd.GetCmd(cfg))
}

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return cfg.Read()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
