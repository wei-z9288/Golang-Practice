package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var versionCmd *cobra.Command

// Define global variables to store the value of flags
var versionFlag string

func init() {
	// Use the StringVarP to bind flag values to variables
	VersionCmd().Flags().StringVarP(&versionFlag, "version", "v", "", "File's version")

	rootcmd.AddCommand(VersionCmd())
}

func SetVersion(newVersion string) {
	Version = newVersion
}

func GetVersion() string {
	return Version
}

func VersionCmd() *cobra.Command {
	if versionCmd == nil {
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Show current version",
			Run: func(cmd *cobra.Command, args []string) {
				// Use the global variable versionFlag
				fmt.Println("fops " + Version + " " + versionFlag)
			},
		}
	}
	return versionCmd
}
