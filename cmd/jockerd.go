/*
Copyright © 2023 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jaronnie/jocker/jockerd"
)

// jockerdCmd represents the jockerd command
var jockerdCmd = &cobra.Command{
	Use:   "jockerd",
	Short: "jockerd daemon",
	Long:  `jockerd daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			jockerd.StartJockerdZrpcServer(cfgFile)
		}()

		go func() {
			jockerd.StartJockerdGatewayServer()
		}()

		select {}
	},
}

func init() {
	rootCmd.AddCommand(jockerdCmd)
}
