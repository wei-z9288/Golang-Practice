package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// Use : The root command of cli. Name of the application
// rootcmd is global var
var rootcmd = &cobra.Command{
	Use:   "fops",
	Short: "Simple CLI for file checking",
	Long: `fops is a simple file check command that checks the number of lines in a file 
	and executes a file checksum operation, it does not support binary files.`,
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		os.Exit(20)
	}
}
