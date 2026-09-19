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
		fmt.Println(version())
	},
}

// version comes from the module version Go stamps into the binary,
// like v0.1.0 for `go install ...@v0.1.0`.
func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		return "dev"
	}
	return info.Main.Version
}

func Cmd() *cobra.Command {
	return versionCmd
}
