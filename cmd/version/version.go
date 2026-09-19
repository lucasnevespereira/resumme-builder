package version

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(String())
	},
}

// String returns the version Go stamps into the binary, like v0.1.0.
func String() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		return "dev"
	}
	return info.Main.Version
}

func Cmd() *cobra.Command {
	return versionCmd
}
