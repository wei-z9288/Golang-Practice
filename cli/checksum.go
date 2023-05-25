package cli

import (
	"fmt"
	algo "fops/algo"
	filesys "fops/filesys"
	"os"

	"github.com/spf13/cobra"
)

var hashFunctions map[string]algo.Algorithm
var checksum *cobra.Command

func init() {

	ChecksumCmd().Flags().StringP("file", "f", "", "File to operate on")

	ChecksumCmd().Flags().Bool("md5", false, "checksum using md5")
	ChecksumCmd().Flags().Bool("sha1", false, "checksum using sha1")
	ChecksumCmd().Flags().Bool("sha256", false, "checksum using sha256")

	ChecksumCmd().MarkFlagRequired("file")

	hashFunctions = map[string]algo.Algorithm{
		"md5":    algo.GetMD5(),
		"sha1":   algo.GetSHA1(),
		"sha256": algo.GetSHA256(),
	}

	rootcmd.AddCommand(ChecksumCmd())
}

func GetChecksum(value []byte, hashFn string) string {
	return hashFunctions[hashFn].GetChecksum(value)
}

func ChecksumRun(cmd *cobra.Command, args []string) error {
	//進行檔案確認
	filename, _ := cmd.Flags().GetString("file")
	fs := filesys.File(filename).CheckFile()

	if fs.Err != nil {
		return fs.Err
	} else {

		algorithmFlag := ""
		contents, _ := os.ReadFile(filename)
		//檢查是下哪個hash flag
		for _, algorithm := range hashFunctions {
			match, err := cmd.Flags().GetBool(algorithm.GetType())
			if err != nil {
				return err
			}
			if match {
				algorithmFlag = algorithm.GetType()
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
