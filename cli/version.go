package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var versionCmd *cobra.Command

// Added the versioncmd command to the rootcmd
func init() {
	VersionCmd().Flags().StringP("version", "v", "", "File's version")
	rootcmd.AddCommand(VersionCmd())
}

func SetVersion(newVersion string) {
	Version = newVersion
}

func GetVersion() string {
	return Version
}

//Run (1)Receive pointer for the command being executed
//    (2)args []string: Parameters input in the command line.
//       This slice contains all parameters except for the command name.

func VersionCmd() *cobra.Command {
	if versionCmd == nil {
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Show current version",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("fops " + Version)
			},
		}
	}
	return versionCmd
}
