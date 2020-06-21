package cmd

import (
	"AA/server1"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateServer1Command())
}

func generateServer1Command() *cobra.Command {

	return &cobra.Command{
		Use:   "start-server1",
		Short: " Please run this command to start server 1",
		Run: func(cmd *cobra.Command, args []string) {
			server1.StartServer1()

		},
	}

}
