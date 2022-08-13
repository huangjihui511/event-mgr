package cmd

import (
	"huangjihui511/event-mgr/pkg/service"

	"github.com/spf13/cobra"
)

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "event service",
	Short: "event service",
	Long:  `event service`,
	Run: func(cmd *cobra.Command, args []string) {
		service.StartService()
	},
}
