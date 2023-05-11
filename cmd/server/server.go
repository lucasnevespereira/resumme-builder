package server

import (
	"github.com/spf13/cobra"
	"resumme-builder/internal/api"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server to use app as an API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := api.New().Run(); err != nil {
			return err
		}
		return nil
	},
}

func Cmd() *cobra.Command {
	return serverCmd
}
