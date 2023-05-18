package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// Use : cli 的根命令。應用程序的名稱
// rootcmd 為global var
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
