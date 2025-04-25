package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootcmd = &cobra.Command{
	Use:   "prometheus-demo",
	Short: "Prometheus-server",
	Long:  `Prometheu Demo Server API ping counter`,
}

func Execute() {
	err := rootcmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
