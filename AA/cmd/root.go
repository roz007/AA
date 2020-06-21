package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	//Use:   "UDP connection scripts",
	//Short: "This package contains scripts helping us route UDP packets sent by Brokers to our remote VPC to local machines",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}



func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
