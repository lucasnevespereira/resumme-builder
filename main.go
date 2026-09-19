package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/lucasnevespereira/resb/cmd/local"
	"github.com/lucasnevespereira/resb/cmd/server"
	"github.com/lucasnevespereira/resb/cmd/version"
	"github.com/lucasnevespereira/resb/internal/utils/logger"
)

var rootCmd = &cobra.Command{
	Use:   "resb",
	Short: "Build a resume PDF from JSON Resume data",
}

func init() {
	rootCmd.AddCommand(local.Cmd())
	rootCmd.AddCommand(server.Cmd())
	rootCmd.AddCommand(version.Cmd())

	// Gives the root command -v and --version, printing the same as `version`.
	rootCmd.Version = version.String()
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.InitDefaultVersionFlag()
}

func main() {
	if len(os.Args) == 1 {
		rootCmd.Help()
		return
	}

	// default cmd if no cmd is given
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && !errors.Is(cmd.Flags().Parse(os.Args[1:]), pflag.ErrHelp) && !cmd.Flags().Changed("version") {
		args := append([]string{local.Cmd().Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatal(err)
	}
}
