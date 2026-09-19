package server

import (
	"github.com/lucasnevespereira/resb/internal/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server to use app as an API",
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := api.New()
		if err != nil {
			return err
		}
		return server.Run()
	},
}

func Cmd() *cobra.Command {
	return serverCmd
}
