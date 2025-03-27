package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

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
	if len(os.Args) == 1 {
		rootCmd.Help()
		return
	}

	// default cmd if no cmd is given
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && !errors.Is(cmd.Flags().Parse(os.Args[1:]), pflag.ErrHelp) {
		args := append([]string{local.Cmd().Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatal(err)
	}
}
