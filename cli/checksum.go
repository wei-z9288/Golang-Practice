package cli

import (
	"fmt"
	algo "fops/algo"
	filesys "fops/filesys"
	"os"

	"github.com/spf13/cobra"
)

var hashFunctions map[string]string
var checksum *cobra.Command

func init() {

	ChecksumCmd().Flags().StringP("file", "f", "", "File to operate on")

	ChecksumCmd().Flags().Bool("md5", false, "checksum using md5")
	ChecksumCmd().Flags().Bool("sha1", false, "checksum using sha1")
	ChecksumCmd().Flags().Bool("sha256", false, "checksum using sha256")

	ChecksumCmd().MarkFlagRequired("file")

	hashFunctions = map[string]string{
		"md5":    "md5",
		"sha1":   "sha1",
		"sha256": "sha256",
	}

	rootcmd.AddCommand(ChecksumCmd())
}

func GetChecksum(value []byte, hashFn string) string {
	return algo.GetChecksum(value, hashFn)
}

func ChecksumRun(cmd *cobra.Command, args []string) error {
	//Verify the file
	filename, _ := cmd.Flags().GetString("file")
	fs := filesys.File(filename).CheckFile()

	if fs.Err != nil {
		return fs.Err
	} else {

		algorithmFlag := ""
		contents, _ := os.ReadFile(filename)
		//Check which hash flag is placed
		for _, hashname := range hashFunctions {
			match, err := cmd.Flags().GetBool(hashname)
			if err != nil {
				return err
			}
			if match {
				algorithmFlag = hashname
			}
		}

		if algorithmFlag == "" {
			return fmt.Errorf("checksum flag is missing. one of md5, sha1, sha256 is required")
		}

		fmt.Println(GetChecksum(contents, algorithmFlag))
	}

	return nil
}

func ChecksumCmd() *cobra.Command {
	if checksum == nil {
		checksum = &cobra.Command{
			Use:   "checksum",
			Short: "Generate checksum of a file",
			Long: `Checksum retrieves the cryptographic hash of a file.
        A valid text file and an algorithm flag (MD5, SHA1, SHA256) are needed.`,
			RunE: ChecksumRun,
		}
	}
	return checksum
}
