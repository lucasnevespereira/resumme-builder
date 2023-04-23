package main

import (
	"github.com/spf13/cobra"
	"resumme-builder/cmd/local"
	"resumme-builder/cmd/server"
	"resumme-builder/internal/utils/logger"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(local.Cmd())
	rootCmd.AddCommand(server.Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatal(err)
	}
}
