package main

import (
	"huangjihui511/event-mgr/cmd"
	"huangjihui511/event-mgr/pkg/logs"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logs.Logger.Error(err)
		os.Exit(1)
	}
}
