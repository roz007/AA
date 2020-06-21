package cmd

import (
	server2 "AA/server2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateServer2Command())
}

func generateServer2Command() *cobra.Command {

	return &cobra.Command{
		Use:   "start-server2",
		Short: " Please run this command to start server 1",
		Run: func(cmd *cobra.Command, args []string) {
			server2.StartServer2()
		},
	}

}
