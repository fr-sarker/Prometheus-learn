package cmd

import (
	"fmt"
	"github.com/frsarker/Prometheus-learn/api"
	"github.com/spf13/cobra"
)

var Port string

var (
	startcCmd = &cobra.Command{
		Use:   "start",
		Short: "start cmd starts the server on a port",
		Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting server on port:", Port)
			api.RunServer(Port)
		},
	}
)

func init() {
	startcCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8080", "Port to run server")
	rootcmd.AddCommand(startcCmd)
}
