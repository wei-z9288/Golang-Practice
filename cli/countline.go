package cli

import (
	"fmt"
	filesys "fops/filesys"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var linecount *cobra.Command

// Define global variables to store the value of flags
var file string

func init() {
	// Use the StringVarP to bind flag values to variables
	LinecountCmd().Flags().StringVarP(&file, "file", "f", "", "File to operate on")

	LinecountCmd().MarkFlagRequired("file")
	rootcmd.AddCommand(LinecountCmd())
}

func count(filename string) (int, error) {

	lines, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error: Couldn't read file '%v'", filename)
	}
	return strings.Count(string(lines), "\n"), err
}

func LinecountRun(cmd *cobra.Command, args []string) error {

	//Conduct a series of file checks
	fs := filesys.File(file).CheckFile().CheckNotBinary()

	if fs.Err != nil {
		fmt.Println(fs.Err)
	} else {
		lines, err := count(fs.Filename)
		if err != nil {
			fmt.Println(err)
			return err
		} else {
			fmt.Println(lines)
		}
	}
	return nil
}

func LinecountCmd() *cobra.Command {
	if linecount == nil {
		linecount = &cobra.Command{
			Use:   "linecount",
			Short: "Count lines in a file",
			Long:  `Linecount counts the number of lines in a text file. A valid text file is required.`,
			RunE:  LinecountRun,
		}
	}
	return linecount
}
